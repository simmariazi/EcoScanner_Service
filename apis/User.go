package apis

import (
	"main/function"

	"github.com/gin-gonic/gin"
)

// get godoc
// @tags customer
// @summary 단일 구매자 조회
// @Param Authorization-Customer header string true "customer-auth-token"
// @Accept  json
// @Produce json
// @Router	/apis/commerce/customer [get]
// @Success 200 "Success"
// @Failure 400 "Validation error"
// @Failure 401 "Authorization error"
// @Failure 404 "Not_Found error"
// @Failure 500 "Application_Server error"
func GetMembers(c *gin.Context) {
	c.JSON(200, function.GetMemberList())
}
