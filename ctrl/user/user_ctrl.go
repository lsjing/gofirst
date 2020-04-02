package user

import (
	//"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lsjing/gofirst/models/users"
	"github.com/lsjing/gofirst/entitys"
	//"github.com/lsjing/gofirst/utils"
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
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 1,
			"msg"  : "name长度错误",
		})
		return
	}

	if ageStrCount > 2 || ageStrCount < 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 1,
			"msg"  : "age长度错误",
		})
		return
	}

	//写入之前查询下是否已存在
	_, has, err := user_model.UserOneByName(name)
	code := 0
	if err != nil {
		code = 1
	}
	if has {
		code = 2
		ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "user exist"})
		return
	}

	ageInt, _ := strconv.Atoi(ageStr)
	res := user_model.UserAdd(&entitys.User{Name: name, Age: ageInt})

	if !res {
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

func GetUserAction(c *gin.Context) {
	uid := c.Param("user_id")

	u, has, err := user_model.UserOneById(uid)
	code := 0
	if err != nil {
		code = 1
	}

	if !has {
		c.JSON(http.StatusOK, gin.H{
			"code"	:	code,
			"msg"	:	"ok",
			"data"	:	"",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code"	:	code,
		"msg"	:	"ok",
		"data"	:	u,
	})

}

func Demo(ctx *gin.Context) {
	u, err := user_model.UserOne()
	code := 0
	if err != false {
		code = 1
	}
	if !u.IsEmpty() {
		ctx.JSON(http.StatusOK, gin.H{
			"code"	:	code,
			"msg"	:	"ok",
			"data"	:	u,
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code":code,"msg":"ok","data":""})
	}

}