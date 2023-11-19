package main

import (
	"github.com/gin-gonic/gin"
	"real-estate-golang-poc.com/V0/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/", controllers.HelloWorld)

	protected := r.Group("/api")
	protected.GET("/ads", controllers.FindAds)

	err := r.Run()
	if err != nil {
		return
	}
}
