package main

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"go-schedule/controllers"
	"go-schedule/scheduler"
)

func main() {
	logger.Println("Server is started ...")

	go scheduler.Run()

	router := gin.Default()

	router.POST("/customers", controllers.Create_customer)
	router.GET("/customers", controllers.Get_all_customer)

	if router.Run("localhost:8080") != nil {
		return
	}
}
