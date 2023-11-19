package main

import (
	"fmt"
	"github.com/MicahParks/keyfunc/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"real-estate-golang-poc.com/V0/controllers"
	"strings"
	"time"
)

func AuthMiddleware(verifier func(token string) bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		fmt.Println(token)

		isValid := verifier(strings.Trim(token, " "))
		log.Printf("iSvalid %v", isValid)
		if !isValid {
			return
		}

		c.Next()
	}
}

func main() {
	jwksURL := "http://localhost:8080/realms/Real-Estate/protocol/openid-connect/certs"

	// Create a context that, when cancelled, ends the JWKS background refresh goroutine.
	//ctx, cancel := context.WithCancel(context.Background())

	// Create the keyfunc options. Use an error handler that logs. Refresh the JWKS when a JWT signed by an unknown KID
	// is found or at the specified interval. Rate limit these refreshes. Timeout the initial JWKS refresh request after
	// 10 seconds. This timeout is also used to create the initial context.Context for keyfunc.Get.
	options := keyfunc.Options{
		//Ctx: ,
		RefreshErrorHandler: func(err error) {
			log.Printf("There was an error with the jwt.Keyfunc\nError: %v", err)
		},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Second * 10,
		RefreshUnknownKID: true,
	}

	// Create the JWKS from the resource at the given URL.
	jwks, err := keyfunc.Get(jwksURL, options)
	if err != nil {
		log.Fatalf("Failed to create JWKS from resource at the given URL.\nError: %v", err)
	}

	verify := func(rawToken string) bool {
		token, err := jwt.Parse(rawToken, jwks.Keyfunc)
		if err != nil {
			log.Printf("Failed to parse the JWT.\nError: %v", err)
			return false
		}
		return token.Valid
	}
	// Parse the JWT.
	//token, err := jwt.Parse(jwtB64, jwks.Keyfunc)
	//if err != nil {
	//	log.Fatalf("Failed to parse the JWT.\nError: %v", err)
	//}
	//
	//// Check if the token is valid.
	//if !token.Valid {
	//	log.Fatalf("The token is not valid.")
	//}
	//log.Println("The token is valid.")

	// End the background refresh goroutine when it's no longer needed.
	//cancel()

	// This will be ineffectual because the line above this canceled the parent context.Context.
	// This method call is idempotent similar to context.CancelFunc.
	//jwks.EndBackground()

	r := gin.Default()

	r.GET("/", controllers.HelloWorld)

	protected := r.Group("/api")
	protected.Use(AuthMiddleware(verify))
	protected.GET("/ads", controllers.FindAds)
	protected.OPTIONS("/ads", controllers.FindAds)

	err = r.Run(":9000")
	if err != nil {
		return
	}
}
