package user

import "github.com/gin-gonic/gin"
import "net/http"

func RunAction(c *gin.Context) {
	c.String(http.StatusOK, "pring info in RunAction,user package,ctrl package")
}