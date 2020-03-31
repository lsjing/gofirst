package router

import (
	"github.com/gin-gonic/gin"
)
import "net/http"
import "github.com/lsjing/gofirst/utils"
import (
	"github.com/lsjing/gofirst/ctrl"
	"github.com/lsjing/gofirst/ctrl/user"
)

var Router *gin.Engine

func init()  {
	Router = gin.Default()
}

func Setup() *gin.Engine  {

	public := utils.Config.Section("router").Key("public").String()
	Router.Static("/public", public)

	Router.GET("/user", func(c *gin.Context) {
		c.String(http.StatusOK, "hello 123")
	})

	Router.GET("/say", ctrl.SayAction)
	Router.GET("/run", user.RunAction)
	Router.GET("/read", user.ReadDBAction)
	Router.POST("/user/add", user.WriteDBAction)

	viewPath := utils.Config.Section("router").Key("view_path").String()
	Router.LoadHTMLGlob(viewPath)

	Router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", "")
	})

	return Router
}
