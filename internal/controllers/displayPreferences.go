package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blyndusk/flamingops/internal/middlewares"
	"github.com/blyndusk/flamingops/pkg/models"
)

func CreateDisplayPreferences(c *gin.Context) {
	var input models.DisplayPreferencesInput
	middlewares.CreateDisplayPreferences(c, &input)
	c.JSON(http.StatusOK, input)
}

func GetDisplayPreferences(c *gin.Context) {
	var preferences models.DisplayPreferences
	middlewares.GetDisplayPreferences(c, &preferences)
	c.JSON(http.StatusOK, preferences)
}

func UpdateDisplayPreferences(c *gin.Context) {
	var preferences models.DisplayPreferences
	var input models.DisplayPreferencesInput
	middlewares.UpdateDisplayPreferences(c, &preferences, &input)
	c.JSON(http.StatusOK, preferences)
}

func DeleteDisplayPreferences(c *gin.Context) {
	var preferences models.DisplayPreferences
	middlewares.DeleteDisplayPreferences(c, &preferences)
	c.JSON(http.StatusOK, gin.H{
		"message": "DisplayPreferences deleted successfully",
	})
}
