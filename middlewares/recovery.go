package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryHandler(c *gin.Context, errs interface{}) {
	if err, ok := errs.(string); ok {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}
