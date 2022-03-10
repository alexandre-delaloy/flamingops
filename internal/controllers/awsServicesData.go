package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blyndusk/flamingops/internal/middlewares"
	"github.com/blyndusk/flamingops/pkg/models"
)

func GetAwsServicesData(c *gin.Context) {
	var awsServicesData models.AwsServicesData
	middlewares.GetAwsServicesData(c, &awsServicesData)
	c.JSON(http.StatusOK, awsServicesData)
}

func DeleteAwsServicesData(c *gin.Context) {
	var awsServicesData models.AwsServicesData
	middlewares.DeleteAwsServicesData(c, &awsServicesData)
	c.JSON(http.StatusOK, gin.H{
		"message": "AwsServicesData deleted successfully",
	})
}
