package middlewares

import (
	"net/http"

	"github.com/blyndusk/flamingops/internal/database"
	"github.com/blyndusk/flamingops/pkg/helpers"
	"github.com/blyndusk/flamingops/pkg/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateActiveServices(c *gin.Context, input *models.ActiveServicesInput) error {
	activeServices := hydrateActiveServices(input)
	if err := database.Db.Create(&activeServices).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func GetActiveServices(c *gin.Context, activeServices *models.ActiveServices) {
	if err := database.Db.Where("user_id = ?", c.GetUint("userId")).First(&activeServices).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func UpdateActiveServices(c *gin.Context, activeServices *models.ActiveServices, input *models.ActiveServicesInput) {
	GetActiveServices(c, activeServices)
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		httpStatus, response := helpers.ErrorToJson(http.StatusBadRequest, err.Error())
		c.JSON(httpStatus, response)
		return
	}

	updatedActiveServices := hydrateActiveServices(input)
	database.Db.Model(&activeServices).Updates(updatedActiveServices)
}

func DeleteActiveServices(c *gin.Context, activeServices *models.ActiveServices) {
	if err := database.Db.Where("user_id = ?", c.GetUint("userId")).First(&activeServices).Delete(&activeServices).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func hydrateActiveServices(input *models.ActiveServicesInput) models.ActiveServices {
	return models.ActiveServices{
		UserId:      input.UserId,
		AwsServices: input.AwsServices,
		SwServices:  input.SwServices,
	}
}
