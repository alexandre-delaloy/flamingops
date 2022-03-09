package middlewares

import (
	"net/http"

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

	CreateActiveServices(c, &models.ActiveServicesInput{UserId: user.ID})
	CreateSwServicesData(c, &models.SwServicesDataInput{UserId: user.ID})
	CreateAwsServicesData(c, &models.AwsServicesDataInput{UserId: user.ID})
	CreateDisplayPreferences(c, &models.DisplayPreferencesInput{UserId: user.ID})
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

	DeleteActiveServices(c, &models.ActiveServices{UserId: c.Params.ByName("id")})
	DeleteSwServicesData(c, &models.SwServicesData{UserId: c.Params.ByName("id")})
	DeleteAwsServicesData(c, &models.AwsServicesData{UserId: c.Params.ByName("id")})
	DeleteDisplayPreferences(c, &models.DisplayPreferences{UserId: c.Params.ByName("id")})
}

func hydrateUser(input *models.UserInput) models.User {
	return models.User{
		Username: input.Username,
		Mail:     input.Mail,
		Password: input.Password,
	}
}
