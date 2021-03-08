package main

import (
	"goshop/src/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/product-type", func(c *gin.Context) {
		controller.CreateProductType(c)
	})

	r.Run(":8000")
}
