package apis

import (
	"main/function"

	"github.com/gin-gonic/gin"
)

// godoc
// @ 멤버 정보 조회
// @Description 멤버 정보 조회
// @name GetMembers
// @Accept  json
// @Produce  json
// @Param void
// @Router /member [get]
// @Success 200 {object} MemberInfo
func GetMembers(c *gin.Context) {
	c.JSON(200, function.GetMemberList())
}
