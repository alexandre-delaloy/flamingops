package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blyndusk/flamingops/internal/middlewares"
	"github.com/blyndusk/flamingops/pkg/models"
)

func GetActiveServices(c *gin.Context) {
	var activeServices models.ActiveServices
	middlewares.GetActiveServices(c, &activeServices)
	c.JSON(http.StatusOK, activeServices)
}

func UpdateActiveServices(c *gin.Context) {
	var activeServices models.ActiveServices
	var input models.ActiveServicesInput
	middlewares.UpdateActiveServices(c, &activeServices, &input)
	c.JSON(http.StatusOK, activeServices)
}

func DeleteActiveServices(c *gin.Context) {
	var activeServices models.ActiveServices
	middlewares.DeleteActiveServices(c, &activeServices)
	c.JSON(http.StatusOK, gin.H{
		"message": "ActiveServices deleted successfully",
	})
}
