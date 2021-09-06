package apis

import (
	"main/function"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	c.JSON(200, function.GetProductSimpleList())
}

func GetWishList(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("memberno"))

	c.JSON(200, function.GetWishList(id))
}

func AddwishListProduct(c *gin.Context) {
	memberNo, _ := strconv.Atoi(c.Query("memberNo"))
	productId, _ := strconv.Atoi(c.Query("productId"))

	c.JSON(200, function.AddWishListProduct(memberNo, productId))
}

func AddwishListSeller(c *gin.Context) {
	memberNo, _ := strconv.Atoi(c.Query("memberNo"))
	sellerId, _ := strconv.Atoi(c.Query("sellerId"))

	c.JSON(200, function.AddwishListSeller(memberNo, sellerId))
}

func DeleteWishList(c *gin.Context) {
	memberNo, _ := strconv.Atoi(c.Param("memberno"))
	itemType := c.Param("type")
	allOrNot, _ := strconv.ParseBool(c.Query("allOrNot"))
	id, _ := strconv.Atoi(c.Query("id"))
	function.DeleteWishList(memberNo, id, itemType, allOrNot)
	c.JSON(200, "삭제 완료")
}

func GetCompareProduct(c *gin.Context) {
	productId := c.Query("productId")

	c.JSON(200, function.GetCompareProduct(productId))
}

func GetProduct(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("productId"))

	c.JSON(200, function.GetProductDetail(productId))
}
