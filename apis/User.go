package apis

import (
	"main/function"

	"github.com/gin-gonic/gin"
)

// GetMembers godoc
// @tags User
// @summary 회원 목록 조회
// @Accept  json
// @Produce json
// @Router	/member [get]
// @Success 200 "Success"
func GetMembers(c *gin.Context) {
	c.JSON(200, function.GetMemberList())
}
