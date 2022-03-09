package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blyndusk/flamingops/internal/middlewares"
	"github.com/blyndusk/flamingops/pkg/models"
)

func CreateUser(c *gin.Context) {
	var input models.UserInput
	middlewares.CreateUser(c, &input)
	c.JSON(http.StatusOK, input)
}

func Login(c *gin.Context) {
	var input models.login
	middlewares.Login(c, &input)
	c.JSON(http.StatusOK, input)
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
	middlewares.DeleteUser(c, &user)
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
