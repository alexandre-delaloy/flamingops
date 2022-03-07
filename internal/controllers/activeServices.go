package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blyndusk/flamingops/internal/middlewares"
	"github.com/blyndusk/flamingops/pkg/models"
)

func CreateActiveServices(c *gin.Context) {
	var input models.ActiveServicesInput
	middlewares.CreateActiveServices(c, &input)
	c.JSON(http.StatusOK, input)
}

func GetActiveServicesById(c *gin.Context) {
	var preferences models.ActiveServices
	middlewares.GetActiveServicesById(c, &preferences)
	c.JSON(http.StatusOK, preferences)
}

func UpdateActiveServices(c *gin.Context) {
	var preferences models.ActiveServices
	var input models.ActiveServicesInput
	middlewares.UpdateActiveServices(c, &preferences, &input)
	c.JSON(http.StatusOK, preferences)
}

func DeleteActiveServices(c *gin.Context) {
	var preferences models.ActiveServices
	middlewares.DeleteActiveServices(c, &preferences)
	c.JSON(http.StatusOK, gin.H{
		"message": "ActiveServices deleted successfully",
	})
}
