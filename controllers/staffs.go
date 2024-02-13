package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-schedule/services/staffs"
	"net/http"
)

func Export_staff(c *gin.Context) {
	errs := staffs.Export_DB()

	if errs != nil {
		fmt.Println(errs.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}
	fmt.Println("Succsessfully Exported Staff DB")
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "Succsessfully Exported Staff DB"})
}
