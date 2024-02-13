package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-schedule/models"
	"go-schedule/services/customers"
	"net/http"
)

func Get_all_customer(c *gin.Context) {
	data, err := customers.Get_all()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": data})
}

func Create_customer(c *gin.Context) {
	var customer models.Customer
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the insert service function
	if err := customers.Create_Customer(customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert customer"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Customer inserted successfully"})
}

func Export_customer(c *gin.Context) {
	err := customers.Export_DB()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "Succsessfully Exported"})
}

func Import_customer(c *gin.Context) {
	err := customers.Import_DB()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "Succsessfully Imported"})
}
