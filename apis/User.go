package apis

import (
	"main/function"

	"github.com/gin-gonic/gin"
)

func GetMembers(c *gin.Context) {
	c.JSON(200, function.GetMemberList())
}
