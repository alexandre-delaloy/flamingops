package middlewares

import (
	"net/http"
	"strconv"

	"github.com/blyndusk/flamingops/internal/database"
	"github.com/blyndusk/flamingops/pkg/helpers"
	"github.com/blyndusk/flamingops/pkg/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateUser(c *gin.Context, input *models.UserInput) {
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		httpStatus, response := helpers.ErrorToJson(http.StatusBadRequest, err.Error())
		c.JSON(httpStatus, response)
		return
	}

	user := hydrateUser(input)
	if err := database.Db.Create(&user).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}

	CreateActiveServices(c, &models.ActiveServicesInput{UserId: user.Id})
	CreateSwServicesData(c, &models.SwServicesDataInput{UserId: user.Id})
	CreateAwsServicesData(c, &models.AwsServicesDataInput{UserId: user.Id})
	CreateRequestedRegions(c, &models.RequestedRegionsInput{UserId: user.Id})
}

func GetAllUsers(c *gin.Context, users *models.Users) {
	if err := database.Db.Find(&users).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func GetUserById(c *gin.Context, user *models.User) {
	if err := database.Db.Where("id = ?", c.Params.ByName("id")).First(&user).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func UpdateUser(c *gin.Context, user *models.User, input *models.UserInput) {
	GetUserById(c, user)
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		httpStatus, response := helpers.ErrorToJson(http.StatusBadRequest, err.Error())
		c.JSON(httpStatus, response)
		return
	}

	updatedUser := hydrateUser(input)
	database.Db.Model(&user).Updates(updatedUser)
}

func DeleteUser(c *gin.Context, user *models.User) {
	if err := database.Db.First(&user, c.Params.ByName("id")).Delete(&user).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}

	i64, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		log.Error(err)
		httpStatus, response := helpers.ErrorToJson(http.StatusBadRequest, err.Error())
		c.JSON(httpStatus, response)
		return
	}

	i := uint(i64)

	DeleteActiveServices(c, &models.ActiveServices{UserId: i})
	DeleteSwServicesData(c, &models.SwServicesData{UserId: i})
	DeleteAwsServicesData(c, &models.AwsServicesData{UserId: i})
	DeleteRequestedRegions(c, &models.RequestedRegions{UserId: i})
}

func hydrateUser(input *models.UserInput) models.User {
	return models.User{
		Username: input.Username,
		Mail:     input.Mail,
		Password: input.Password,
	}
}
