package router

import "github.com/gin-gonic/gin"
import "net/http"
import "github.com/lsjing/gofirst/utils"

var Router *gin.Engine

func init()  {
	Router = gin.Default()
}

func Setup() *gin.Engine  {

	public := utils.Config.Section("router").Key("public").String()
	Router.Static("/public", public)

	viewPath := utils.Config.Section("router").Key("view_path").String()
	Router.LoadHTMLGlob(viewPath)

	Router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", "")
	})

	return Router
}
