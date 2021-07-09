package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/v", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"project": "lkapi-supercharged",
			"author":  "michaelpeterswa",
			"v":       "0.0.1",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
