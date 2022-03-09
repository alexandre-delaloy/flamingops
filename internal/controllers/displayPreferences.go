package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blyndusk/flamingops/internal/middlewares"
	"github.com/blyndusk/flamingops/pkg/models"
)

func CreateRequestedRegions(c *gin.Context) {
	var input models.RequestedRegionsInput
	middlewares.CreateRequestedRegions(c, &input)
	c.JSON(http.StatusOK, input)
}

func GetRequestedRegions(c *gin.Context) {
	var preferences models.RequestedRegions
	middlewares.GetRequestedRegions(c, &preferences)
	c.JSON(http.StatusOK, preferences)
}

func UpdateRequestedRegions(c *gin.Context) {
	var preferences models.RequestedRegions
	var input models.RequestedRegionsInput
	middlewares.UpdateRequestedRegions(c, &preferences, &input)
	c.JSON(http.StatusOK, preferences)
}

func DeleteRequestedRegions(c *gin.Context) {
	var preferences models.RequestedRegions
	middlewares.DeleteRequestedRegions(c, &preferences)
	c.JSON(http.StatusOK, gin.H{
		"message": "RequestedRegions deleted successfully",
	})
}
