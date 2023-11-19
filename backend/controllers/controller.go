package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Ad struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
}

func FindAds(c *gin.Context) {
	ads := []Ad{
		{
			ID:          1,
			Title:       "Super Apartment in Saint Denis",
			Description: "Will make you fall in love with this city",
			Price:       100_000,
		},
		{
			ID:          2,
			Title:       "Apartment to renovate Paris 19",
			Description: "Needs renovation but close proximity from everything you need",
			Price:       400_000,
		},
	}
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Mode, Authorization, User-Agent, Dnt, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform")

	c.JSON(http.StatusOK, gin.H{"ads": ads})
}
