package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lsjing/gofirst/models/users"
	"github.com/lsjing/gofirst/entitys"
	"github.com/lsjing/gofirst/utils"
	"strconv"
	"github.com/thedevsaddam/govalidator"
	"net/http"
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
	rules := govalidator.MapData{
		"name": []string{"required", "between:3,8"},
		"age":  []string{"digits:2"},
	}

	messages := govalidator.MapData{
		"name": []string{"required:用户名不能为空", "between:3到8位"},
		"age":  []string{"digits:手机号码为11位数字"},
	}

	opts := govalidator.Options{
		Request:         ctx.Request, // request object
		Rules:           rules,       // rules map
		Messages:        messages,    // custom message map (Optional)
		RequiredDefault: false,       // all the field to be pass the rules
	}
	v := govalidator.New(opts)
	e := v.Validate()

	//校验结果判断
	if len(e)>0 {
		ctx.JSON(200, e)
		return
	}

	name := ctx.PostForm("name")
	age_str := ctx.PostForm("age")
	age_int, _ := strconv.Atoi(age_str)
	var err bool
	err = user_model.UserAdd(&entitys.User{Name: name, Age: age_int})

	ctx.JSON(http.StatusOK,utils.ResponseNode{Code: 0, Msg:"123", Data:"ssss", Val: err})

}