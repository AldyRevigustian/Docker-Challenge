package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello, from Golang",
		})
	})
	r.GET("/map", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello, from endpoint 1asdsad",
		})
	})

	r.GET("/endpoint2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello, from endpoint 2!",
		})
	})

	r.Run(":8181")
}
