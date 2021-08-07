package main

import (
	function "main/function"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/seller", func(c *gin.Context) {
		c.JSON(200, function.SelectMemeber())
	})
	r.Run()
}
