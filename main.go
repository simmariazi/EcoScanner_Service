package main

import (
	"main/database"
	function "main/function"
	"main/middlewares"
	"os"
	"strconv"

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

	r.GET("/wishlist/:memberno", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("memberno"))

		c.JSON(200, function.GetWishList(id))
	})

	r.PUT("/wishlist/product", func(c *gin.Context) {
		memberNo, _ := strconv.Atoi(c.Query("memberNo"))
		productId, _ := strconv.Atoi(c.Query("productId"))

		c.JSON(200, function.AddWishListProduct(memberNo, productId))
	})

	r.PUT("/wishlist/seller", func(c *gin.Context) {
		memberNo, _ := strconv.Atoi(c.Query("memberNo"))
		sellerId, _ := strconv.Atoi(c.Query("sellerId"))

		c.JSON(200, function.AddwishListSeller(memberNo, sellerId))
	})

	r.DELETE("/wishlist/:type/:memberno", func(c *gin.Context) {
		memberNo, _ := strconv.Atoi(c.Param("memberno"))
		itemType := c.Param("type")
		allOrNot, _ := strconv.ParseBool(c.Query("allOrNot"))
		id, _ := strconv.Atoi(c.Query("id"))
		function.DeleteWishList(memberNo, id, itemType, allOrNot)
		c.JSON(200, "삭제 완료")
	})

	r.GET("/product/compare", func(c *gin.Context) {
		productId := c.Query("productId")

		c.JSON(200, function.GetCompareProduct(productId))
	})

	r.GET("/board", func(c *gin.Context) {
		c.JSON(200, function.GetBoardList())
	})

	useApp(r)

	r.Run(":" + os.Getenv("PORT"))
}

func useApp(app *gin.Engine) {
	app.Use(middlewares.CORSMiddleware())
	app.Use(gin.Logger())
	app.Use(gin.CustomRecovery(middlewares.RecoveryHandler))
}
