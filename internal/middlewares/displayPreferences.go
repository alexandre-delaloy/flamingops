package middlewares

import (
	"net/http"

	"github.com/blyndusk/flamingops/internal/database"
	"github.com/blyndusk/flamingops/pkg/helpers"
	"github.com/blyndusk/flamingops/pkg/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateDisplayPreferences(c *gin.Context, input *models.DisplayPreferencesInput) {
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		httpStatus, response := helpers.ErrorToJson(http.StatusBadRequest, err.Error())
		c.JSON(httpStatus, response)
		return
	}

	displayPreferences := hydrateDisplayPreferences(input)
	if err := database.Db.Create(&displayPreferences).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func GetDisplayPreferences(c *gin.Context, displayPreferences *models.DisplayPreferences) {
	if err := database.Db.Where("id = ?", c.Params.ByName("id")).First(&displayPreferences).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func UpdateDisplayPreferences(c *gin.Context, displayPreferences *models.DisplayPreferences, input *models.DisplayPreferencesInput) {
	GetDisplayPreferences(c, displayPreferences)
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		httpStatus, response := helpers.ErrorToJson(http.StatusBadRequest, err.Error())
		c.JSON(httpStatus, response)
		return
	}

	updatedDisplayPreferences := hydrateDisplayPreferences(input)
	database.Db.Model(&displayPreferences).Updates(updatedDisplayPreferences)
}

func DeleteDisplayPreferences(c *gin.Context, displayPreferences *models.DisplayPreferences) {
	if err := database.Db.Where("id = ?", c.Params.ByName("id")).First(&displayPreferences).Delete(&displayPreferences).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func hydrateDisplayPreferences(input *models.DisplayPreferencesInput) models.DisplayPreferences {
	return models.DisplayPreferences{
		UserId:  input.UserId,
		Content: input.Content,
	}
}
