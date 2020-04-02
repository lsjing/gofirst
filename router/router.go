package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/lsjing/gofirst/utils"
	"github.com/lsjing/gofirst/ctrl"
	"github.com/lsjing/gofirst/ctrl/user"
)

var Router *gin.Engine

func init()  {
	Router = gin.Default()
}

func Setup() *gin.Engine {

	public := utils.Config.Section("router").Key("public").String()
	Router.Static("/public", public)

	Router.GET("/user", func(c *gin.Context) {
		c.String(http.StatusOK, "hello 123")
	})

	Router.GET("/say", ctrl.SayAction)
	Router.GET("/run", user.RunAction)
	Router.GET("/read", user.ReadDBAction)
	Router.POST("/user/add", user.WriteDBAction)
	//Router.GET("/show", user.UserShowAction)
	Router.GET("/list", user.UserListAction)

	Router.GET("/get/:user_id", user.GetUserAction)
	viewPath := utils.Config.Section("router").Key("view_path").String()
	Router.LoadHTMLGlob(viewPath)

	Router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", "")
	})

	return Router
}
