package main

import (
	function "main/function"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/member", func(c *gin.Context) {
		c.JSON(200, function.SelectMember())
	})

	r.GET("/seller", func(c *gin.Context) {
		c.JSON(200, function.SelectSeller())
	})

	r.Run()
}
