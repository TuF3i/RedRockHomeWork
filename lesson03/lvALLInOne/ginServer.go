package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Bob struct {
	Name   string    `json:"name"`
	Scores []float64 `json:"score"`
}

func HelloHandleFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Query("msg")

		if query == "ping" {
			c.JSON(200, gin.H{"data": "pong"})
			return
		}

		if query == "helloserver" {
			c.JSON(200, gin.H{"data": "helloclient"})
			return
		}
	}
}

func ScoreHandleFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		var scoreRevData Bob

		c.ShouldBindJSON(&scoreRevData)

		var sum float64
		for _, score := range scoreRevData.Scores {
			sum += score
		}

		ave := sum / float64(len(scoreRevData.Scores))

		c.JSON(200, gin.H{"average": ave})

	}
}

func main() {

	fmt.Println("[GET]Display Cat.jpg:       /pic/cat.jpg")
	fmt.Println("[GET]RedRock Page:          /redrock")
	fmt.Println("[GET]Talk to Server:        /talk")
	fmt.Println("[POST]Score Ave Calculator: /score")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Static("/pic", "./data")
	r.StaticFile("/redrock", "data/index.html")

	r.GET("/talk", HelloHandleFunc())
	r.POST("/score", ScoreHandleFunc())

	r.Run(":8080")
}
