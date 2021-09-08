package apis

import (
	"main/function"

	"github.com/gin-gonic/gin"
)

// GetSellers godoc
// @tags Seller
// @summary 기업 목록 조회
// @Accept  json
// @Produce json
// @Router	/seller [get]
// @Success 200 "Success"
func GetSellers(c *gin.Context) {
	c.JSON(200, function.GetSellerList())
}
