package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()


	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message" : "OK!!",
		})
	})


	server.Run(":8080")
}