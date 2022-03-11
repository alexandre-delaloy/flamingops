package middlewares

import (

	"github.com/blyndusk/flamingops/internal/database"
	"github.com/blyndusk/flamingops/pkg/helpers"
	"github.com/blyndusk/flamingops/pkg/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateSwServicesData(c *gin.Context, input *models.SwServicesDataInput) (error) {
	swServicesData := hydrateSwServicesData(input)
	if err := database.Db.Create(&swServicesData).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func GetSwServicesData(c *gin.Context, swServicesData *models.SwServicesData) {
	if err := database.Db.Where("user_id = ?", c.GetUint("userId")).First(&swServicesData).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func DeleteSwServicesData(c *gin.Context, swServicesData *models.SwServicesData) {
	if err := database.Db.Where("user_id = ?", c.GetUint("userId")).First(&swServicesData).Delete(&swServicesData).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func hydrateSwServicesData(input *models.SwServicesDataInput) models.SwServicesData {
	return models.SwServicesData{
		UserId: input.UserId,
		Data:   "",
	}
}
