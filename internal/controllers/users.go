package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/blyndusk/flamingops/internal/middlewares"
	"github.com/blyndusk/flamingops/pkg/helpers"
	"github.com/blyndusk/flamingops/pkg/models"
)

func CreateUser(c *gin.Context) {
	var input models.UserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}

	res, err := middlewares.CreateUser(c, &input)
	if err != nil {
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}

	err = middlewares.CreateActiveServices(c, &models.ActiveServicesInput{UserId: res.Id})
	if err != nil {
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}

	err = middlewares.CreateAwsServicesData(c, &models.AwsServicesDataInput{UserId: res.Id})
	if err != nil {
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}

	err = middlewares.CreateSwServicesData(c, &models.SwServicesDataInput{UserId: res.Id})
	if err != nil {
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}

	err = middlewares.CreateRequestedRegions(c, &models.RequestedRegionsInput{UserId: res.Id})
	if err != nil {
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}

	c.JSON(http.StatusOK, &res)
}

func Login(c *gin.Context) {
	var input models.Login
	res, err := middlewares.Login(c, &input)
	if err != nil {
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
	c.JSON(http.StatusOK, &res)
}

func GetAllUsers(c *gin.Context) {
	var users models.Users
	middlewares.GetAllUsers(c, &users)
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	var user models.User
	middlewares.GetUserById(c, &user)
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	var input models.UserInput
	middlewares.UpdateUser(c, &user, &input)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User

	if err := middlewares.DeleteUser(c, &user); err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
