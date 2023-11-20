package main

import (
	"fmt"
	"github.com/MicahParks/keyfunc/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"real-estate-golang-poc.com/V0/controllers"
	"strings"
	"time"
)

type MyCustomClaims struct {
	Scope string `json:"scope"`
	jwt.RegisteredClaims
}

func AccessAdsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		content, exists := c.Get("claims")
		if !exists {
			fmt.Println("Oops, context wasn't properly set")
			return
		}

		claims, valid := content.(MyCustomClaims)
		if !valid {
			fmt.Println("Oops, token isn't good")
			return
		}
		scope := claims.Scope
		fmt.Println(scope)
		if !strings.Contains(scope, "read:ads") {
			fmt.Println("Oops, forbidden")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func AuthMiddleware(verifier func(token string) (*jwt.Token, MyCustomClaims, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {

			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Mode, Authorization, User-Agent, Dnt, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform")
			c.AbortWithStatus(http.StatusOK)
			return
		}
		rawToken := c.Request.Header.Get("Authorization")
		fmt.Println(rawToken)

		token, claims, err := verifier(strings.Trim(rawToken, " "))
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		log.Printf("iSvalid %v", token.Valid)
		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("token", token)
		c.Set("claims", claims)
		c.Next()
	}
}

func main() {
	jwksURL := "http://localhost:8080/realms/Real-Estate/protocol/openid-connect/certs"

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

	verify := func(rawToken string) (*jwt.Token, MyCustomClaims, error) {
		var customClaims MyCustomClaims
		token, err := jwt.ParseWithClaims(rawToken, &customClaims, jwks.Keyfunc)
		return token, customClaims, err
	}

	r := gin.Default()

	r.GET("/", controllers.HelloWorld)

	protected := r.Group("/api")
	protected.Use(AuthMiddleware(verify))
	protected.Use(AccessAdsMiddleware())
	protected.GET("/ads", controllers.FindAds)
	protected.OPTIONS("/ads", controllers.FindAds)

	err = r.Run(":9000")
	if err != nil {
		return
	}
}
