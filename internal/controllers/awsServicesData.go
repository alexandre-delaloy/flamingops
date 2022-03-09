package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blyndusk/flamingops/internal/middlewares"
	"github.com/blyndusk/flamingops/pkg/models"
)

func CreateAwqServicesData(c *gin.Context) {
	var input models.AwqServicesDataInput
	middlewares.CreateAwqServicesData(c, &input)
	c.JSON(http.StatusOK, input)
}

func GetAwqServicesData(c *gin.Context) {
	var awsServicesData models.AwqServicesData
	middlewares.GetAwqServicesData(c, &awsServicesData)
	c.JSON(http.StatusOK, awsServicesData)
}

func DeleteAwqServicesData(c *gin.Context) {
	var awsServicesData models.AwqServicesData
	middlewares.DeleteAwqServicesData(c, &awsServicesData)
	c.JSON(http.StatusOK, gin.H{
		"message": "AwqServicesData deleted successfully",
	})
}
