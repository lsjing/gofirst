package user

import (
	//"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lsjing/gofirst/models/users"
	"github.com/lsjing/gofirst/entitys"
	"github.com/lsjing/gofirst/utils"
	"strconv"
	//"github.com/thedevsaddam/govalidator"
	"net/http"
	"unicode/utf8"
)

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

func WriteDBAction(ctx *gin.Context) {
	//参数检验
	name := ctx.PostForm("name")
	ageStr := ctx.PostForm("age")

	nameCount := utf8.RuneCountInString(name)
	ageStrCount := utf8.RuneCountInString(ageStr)

	if nameCount > 8 || nameCount < 3 {
		ctx.JSON(http.StatusOK, utils.ResponseNode{Code: 1, Msg:"name长度错误"})
		return
	}

	if ageStrCount > 2 || ageStrCount < 1 {
		ctx.JSON(http.StatusOK, utils.ResponseNode{Code: 2, Msg:"age长度错误"})
		return
	}

	age_int, _ := strconv.Atoi(ageStr)
	err := user_model.UserAdd(&entitys.User{Name: name, Age: age_int})

	code := 0
	if !err {
		code = 3
	}

	ctx.JSON(http.StatusOK, gin.H{"code":code, "msg":"ok"})
}

func UserListAction(ctx *gin.Context) {
	xusers, err := user_model.UserList()
	code := 0
	if err != nil {
		code = 1
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg" : "ok",
		"data": xusers,
	})
}