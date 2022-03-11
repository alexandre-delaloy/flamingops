package middlewares

import (
	"errors"

	"github.com/blyndusk/flamingops/internal/database"
	"github.com/blyndusk/flamingops/internal/utils"
	"github.com/blyndusk/flamingops/pkg/helpers"
	"github.com/blyndusk/flamingops/pkg/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context, input *models.UserInput) (*models.User, error) {
	// Check if user already exists
	var mailCheck models.User
	database.Db.Where("Mail = ?", &input.Mail).First(&mailCheck)
	if mailCheck.Id != 0 {
		log.Error("e-mail already used")
		return nil, errors.New("e-mail already used")
	}

	user := hydrateUser(input)
	if err := database.Db.Create(&user).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return &user, nil
}

func Login(c *gin.Context, input *models.Login) (string, error) {

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		return "", err
	}

	var user models.User
	if err := database.Db.Where("Mail = ?", &input.Mail).First(&user).Error; err != nil {
		log.Error(err)
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		log.Error("wrong password")
		return "", errors.New("passwords don't match")
	}

	jwtToken := utils.GenerateToken(user.Id)

	return jwtToken, nil
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
	if err := database.Db.Where("id = ?", c.GetUint("userId")).First(&user).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func UpdateUser(c *gin.Context, user *models.User, input *models.UserInput) error {
	GetUserById(c, user)

	updatedUser := hydrateUser(input)
	if err := database.Db.Model(&user).Updates(updatedUser).Error; err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func DeleteUser(c *gin.Context, user *models.User) error {

	i := c.GetUint("userId")

	if err := database.Db.First(&user, &i).First(&user).Error; err != nil {
		log.Error(err)
		return err
	}

	DeleteActiveServices(c, &models.ActiveServices{UserId: i})
	DeleteSwServicesData(c, &models.SwServicesData{UserId: i})
	DeleteAwsServicesData(c, &models.AwsServicesData{UserId: i})
	DeleteRequestedRegions(c, &models.RequestedRegions{UserId: i})

	if err := database.Db.First(&user, c.GetUint("userId")).Delete(&user).Error; err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func hydrateUser(input *models.UserInput) models.User {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	return models.User{
		Username: input.Username,
		Mail:     input.Mail,
		Password: string(hashed),
	}
}
