package apis

import (
	"main/function"

	"github.com/gin-gonic/gin"
)

func GetSellers(c *gin.Context) {
	c.JSON(200, function.GetSellerList())
}
