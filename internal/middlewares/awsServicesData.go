package middlewares

import (
	"net/http"

	"github.com/blyndusk/flamingops/internal/database"
	"github.com/blyndusk/flamingops/pkg/helpers"
	"github.com/blyndusk/flamingops/pkg/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateAwsServicesData(c *gin.Context, input *models.AwsServicesDataInput) {
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		httpStatus, response := helpers.ErrorToJson(http.StatusBadRequest, err.Error())
		c.JSON(httpStatus, response)
		return
	}

	awsServicesData := hydrateAwsServicesData(input)
	if err := database.Db.Create(&awsServicesData).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func GetAwsServicesData(c *gin.Context, awsServicesData *models.AwsServicesData) {
	if err := database.Db.Where("user_id = ?", c.Params.ByName("user_id")).First(&awsServicesData).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func DeleteAwsServicesData(c *gin.Context, awsServicesData *models.AwsServicesData) {
	if err := database.Db.Where("user_id = ?", c.Params.ByName("user_id")).First(&awsServicesData).Delete(&awsServicesData).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func hydrateAwsServicesData(input *models.AwsServicesDataInput) models.AwsServicesData {
	return models.AwsServicesData{
		UserId: input.UserId,
		Data:   "",
	}
}
