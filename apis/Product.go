package apis

import (
	"main/function"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetProducts godoc
// @tags Product
// @summary 상품 목록 조회
// @Accept  json
// @Produce json
// @Router	/productsimple [get]
// @Success 200 "Success"
func GetProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	c.JSON(200, function.GetProductSimpleList(page))

}

// GetWishList godoc
// @tags Product
// @summary 찜 목록 조회
// @Param memberNo path string true "회원 번호"
// @Accept  json
// @Produce json
// @Router	/wishlist/{memberNo} [get]
// @Success 200 "Success"
func GetWishList(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("memberno"))

	c.JSON(200, function.GetWishList(id))
}

// AddWishListProduct godoc
// @tags Product
// @summary 찜 상품 추가
// @Param memberNo query string true "회원 번호"
// @Param productId query string true "SKU"
// @Accept  json
// @Produce json
// @Router	/wishlist/product [put]
// @Success 200 "Success"
func AddwishListProduct(c *gin.Context) {
	memberNo, _ := strconv.Atoi(c.Query("memberNo"))
	productId, _ := strconv.Atoi(c.Query("productId"))

	c.JSON(200, function.AddWishListProduct(memberNo, productId))
}

// AddWishListSeller godoc
// @tags Product
// @summary 찜 기업 추가
// @Param memberNo query string true "회원 번호"
// @Param sellerId query string true "기업 번호"
// @Accept  json
// @Produce json
// @Router	/wishlist/seller [put]
// @Success 200 "Success"
func AddwishListSeller(c *gin.Context) {
	memberNo, _ := strconv.Atoi(c.Query("memberNo"))
	sellerId, _ := strconv.Atoi(c.Query("sellerId"))

	c.JSON(200, function.AddwishListSeller(memberNo, sellerId))
}

// DeleteWishList godoc
// @tags Product
// @summary 찜 목록 삭제
// @Param memberNo path string true "회원 번호"
// @Param itemtype path string true "상품 or 기업 선택"
// @Param allOrNot query string true "전체 삭제 or 특정 항목 삭제"
// @Param id query int true "SKU or 기업 번호"
// @Accept  json
// @Produce json
// @Router	/wishlist/{type}/{memberno} [delete]
// @Success 200 "Success"
func DeleteWishList(c *gin.Context) {
	memberNo, _ := strconv.Atoi(c.Param("memberno"))
	itemType := c.Param("type")
	allOrNot, _ := strconv.ParseBool(c.Query("allOrNot"))
	id, _ := strconv.Atoi(c.Query("id"))
	function.DeleteWishList(memberNo, id, itemType, allOrNot)
	c.JSON(200, "삭제 완료")
}

// GetCompareProduct godoc
// @tags Product
// @summary 상품 비교
// @Param productId query string true "SKU"
// @Accept  json
// @Produce json
// @Router	/product/compare [get]
// @Success 200 "Success"
func GetCompareProduct(c *gin.Context) {
	productId := c.Query("productId")

	c.JSON(200, function.GetCompareProduct(productId))
}

// GetProduct godoc
// @tags Product
// @summary 상품 상세 조회
// @Param productId path string true "SKU"
// @Accept  json
// @Produce json
// @Router	/product/{productId} [get]
// @Success 200 "Success"
func GetProduct(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("productId"))

	c.JSON(200, function.GetProductDetail(productId))
}
