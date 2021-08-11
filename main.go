package main

import (
	"main/database"
	function "main/function"
	"main/middlewares"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	database.InitConnectionString(os.Getenv("CONNECTION_STRING"))

	r := gin.Default()
	r.GET("/member", func(c *gin.Context) {
		c.JSON(200, function.GetMemberList())
	})

	r.GET("/seller", func(c *gin.Context) {
		c.JSON(200, function.GetSellerList())
	})

	r.GET("/productsimple", func(c *gin.Context) {
		c.JSON(200, function.GetProductSimpleList())
	})

	useApp(r)

	r.Run(":" + os.Getenv("PORT"))
}

func useApp(app *gin.Engine) {
	app.Use(middlewares.CORSMiddleware())
	app.Use(gin.Logger())
	app.Use(gin.CustomRecovery(middlewares.RecoveryHandler))
}
