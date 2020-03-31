package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lsjing/gofirst/models/users"
)
import "net/http"

func RunAction(c *gin.Context) {
	c.String(http.StatusOK, "pring info in RunAction,user package,ctrl package")
}

func ReadDBAction(ctx *gin.Context) {

	users, err := user_model.UserList()

	if err == nil {
		for _, user := range users {

			fmt.Println(user.Name)
		}

	}
}