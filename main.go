package main

import "github.com/lsjing/gofirst/router"
import "github.com/lsjing/gofirst/utils"
import "github.com/gin-gonic/gin"
import "github.com/robfig/cron"
import "github.com/lsjing/gofirst/models"

func main() {
	gin.SetMode(gin.DebugMode)
	r := router.Setup()

	c := cron.New()
	c.AddFunc("*/60 * * * * *", models.DbCheck)
	c.Start()

	port := utils.Config.Section("system").Key("http_port").String()
	r.Run(":" + port)
}
