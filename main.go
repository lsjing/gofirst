package main

import "github.com/lsjing/gofirst/router"
import "github.com/lsjing/gofirst/utils"
import "github.com/gin-gonic/gin"

func main() {
	gin.SetMode(gin.DebugMode)
	r := router.Setup()
	port := utils.Config.Section("system").Key("http_port").String()
	r.Run(":" + port)
}
