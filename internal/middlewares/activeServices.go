package middlewares

import (
	"net/http"

	"github.com/blyndusk/flamingops/internal/database"
	"github.com/blyndusk/flamingops/pkg/helpers"
	"github.com/blyndusk/flamingops/pkg/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateActiveServices(c *gin.Context, input *models.ActiveServicesInput) {
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		httpStatus, response := helpers.ErrorToJson(http.StatusBadRequest, err.Error())
		c.JSON(httpStatus, response)
		return
	}

	displayPreferences := hydrateActiveServices(input)
	if err := database.Db.Create(&displayPreferences).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func GetAllActiveServices(c *gin.Context, displayPreferences *models.ActiveServices) {
	if err := database.Db.Find(&displayPreferences).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func GetActiveServicesById(c *gin.Context, displayPreferences *models.ActiveServices) {
	if err := database.Db.Where("id = ?", c.Params.ByName("id")).First(&displayPreferences).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func UpdateActiveServices(c *gin.Context, displayPreferences *models.ActiveServices, input *models.ActiveServicesInput) {
	GetActiveServicesById(c, displayPreferences)
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		httpStatus, response := helpers.ErrorToJson(http.StatusBadRequest, err.Error())
		c.JSON(httpStatus, response)
		return
	}

	updatedActiveServices := hydrateActiveServices(input)
	database.Db.Model(&displayPreferences).Updates(updatedActiveServices)
}

func DeleteActiveServices(c *gin.Context, displayPreferences *models.ActiveServices) {
	if err := database.Db.First(&displayPreferences, c.Params.ByName("id")).Delete(&displayPreferences).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func hydrateActiveServices(input *models.ActiveServicesInput) models.ActiveServices {
	return models.ActiveServices{
		ActiveServicesname: input.ActiveServicesname,
		Mail:  input.Mail,
		Password:  input.Password,
	}
}
