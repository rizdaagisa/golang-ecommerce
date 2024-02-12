package controllers

import (
	"github.com/gin-gonic/gin"
	"go-schedule/models"
	customers "go-schedule/services"
	"net/http"
)

func Get_all_customer(c *gin.Context) {
	data := customers.Get_all()
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
