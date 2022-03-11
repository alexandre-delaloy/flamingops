package middlewares

import (
	"net/http"

	"github.com/blyndusk/flamingops/internal/database"
	"github.com/blyndusk/flamingops/pkg/helpers"
	"github.com/blyndusk/flamingops/pkg/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateRequestedRegions(c *gin.Context, input *models.RequestedRegionsInput) error {
	requestedRegions := hydrateRequestedRegions(input)
	if err := database.Db.Create(&requestedRegions).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func GetRequestedRegions(c *gin.Context, requestedRegions *models.RequestedRegions) {
	if err := database.Db.Where("user_id = ?", c.GetUint("userId")).First(&requestedRegions).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func UpdateRequestedRegions(c *gin.Context, requestedRegions *models.RequestedRegions, input *models.RequestedRegionsInput) {
	GetRequestedRegions(c, requestedRegions)
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		httpStatus, response := helpers.ErrorToJson(http.StatusBadRequest, err.Error())
		c.JSON(httpStatus, response)
		return
	}

	updatedRequestedRegions := hydrateRequestedRegions(input)
	database.Db.Model(&requestedRegions).Updates(updatedRequestedRegions)
}

func DeleteRequestedRegions(c *gin.Context, requestedRegions *models.RequestedRegions) {
	if err := database.Db.Where("user_id = ?", c.GetUint("userId")).First(&requestedRegions).Delete(&requestedRegions).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func hydrateRequestedRegions(input *models.RequestedRegionsInput) models.RequestedRegions {
	return models.RequestedRegions{
		UserId:    input.UserId,
		AwsRegion: input.AwsRegion,
		SwRegion:  input.SwRegion,
	}
}
