package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SayAction(c *gin.Context) {
	c.String(http.StatusOK, "print msg in SayAction/ctrl")
}