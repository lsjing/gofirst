package main

import "github.com/lsjing/gofirst/router"
import "github.com/gin-gonic/gin"
import "github.com/lsjing/gofirst/utils"

func main(){
	gin.SetMode(gin.DebugMode)

	r := router.Setup()

	port := utils.Config.Section("system").Key("http_port").String()
	r.Run(":" + port)
}