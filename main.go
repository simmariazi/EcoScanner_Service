package main

import (
	apis "main/apis"
	"main/database"
	"main/middlewares"
	"os"

	_ "main/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title EcoScanner API Docs with Swagger
// @version 1.0
// @host 121.166.4.186:8128
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	database.InitConnectionString(os.Getenv("CONNECTION_STRING"))

	r := gin.Default()

	// User
	r.GET("user/member", apis.GetMembers)

	// Seller
	r.GET("/seller/get", apis.GetSellers)

	// Product
	r.GET("/productsimple", apis.GetProducts)

	r.GET("/product/:productId", apis.GetProduct)

	r.GET("/wishlist/:memberno", apis.GetWishList)

	r.PUT("/wishlist/product", apis.AddwishListProduct)

	r.PUT("/wishlist/seller", apis.AddwishListSeller)

	r.DELETE("/wishlist/:type/:memberno", apis.DeleteWishList)

	r.GET("/product/compare", apis.GetCompareProduct)

	r.GET("/product/search/:productName", apis.Searchproducts)

	// Board
	r.GET("/board", apis.GetBoards)

	r.PUT("/board/post", apis.WriteBoard)

	r.GET("/board/post", apis.GetBoard)

	r.POST("/board/post", apis.ModifyBoard)

	r.DELETE("/board", apis.DeleteBoard)

	//Review
	r.GET("/review", apis.GetReviews)

	r.PUT("/review/post", apis.WriteReview)

	r.POST("/review/post", apis.ModifyReview)

	r.DELETE("/review", apis.DeleteReview)

	useApp(r)

	r.Run(":" + os.Getenv("PORT"))
}

func useApp(app *gin.Engine) {
	app.Use(middlewares.CORSMiddleware())
	app.Use(gin.Logger())
	app.Use(gin.CustomRecovery(middlewares.RecoveryHandler))
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
