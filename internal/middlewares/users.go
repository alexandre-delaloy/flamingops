package middlewares

import (
	"net/http"

	"github.com/blyndusk/flamingops/internal/database"
	"github.com/blyndusk/flamingops/pkg/helpers"
	"github.com/blyndusk/flamingops/pkg/models"
	"github.com/blyndusk/flamingops/pkg/forms"
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
}

func Login(c *gin.Context, input *models.Login) {

	var loginForm forms.LoginForm
	if err:= json.NewDecoder(r.Body).Decode(&loginForm); err != nil {
		log.Println("cannot decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.Model(&models.User{}).Where("username = ?", loginForm.Username).Take(&user).Error; err != nil {
		log.Println("user not found: " + loginForm.Username)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if err := user.CheckPassword(loginForm.Password); err != nil {
		log.Println("wrong password for: " + loginForm.Username)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	jwtToken := utils.GenerateToken(user.ID)

	if err:=json.NewEncoder(w).Encode(jwtToken); err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
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
}

func hydrateUser(input *models.UserInput) models.User {
	return models.User{
		Username: input.Username,
		Mail:  input.Mail,
		Password:  input.Password,
	}
}
