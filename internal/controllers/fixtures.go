package controllers

import (
	"net/http"

	"github.com/blyndusk/flamingops/internal/database"
	"github.com/blyndusk/flamingops/pkg/models"
	"github.com/gin-gonic/gin"
)

func LoadData(c *gin.Context) {
	user := models.User{
		Username: "Alex",
		Mail:  "contact@flamingops.com",
		Password:  "azertyuiop",
	}
	database.Db.Create(&user)
	c.JSON(http.StatusOK, user)
}
