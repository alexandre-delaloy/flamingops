package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/blyndusk/flamingops/internal/database"
	"github.com/blyndusk/flamingops/internal/queue"
	"github.com/blyndusk/flamingops/pkg/models"
	log "github.com/sirupsen/logrus"
)

func ManuallySendMessage(c *gin.Context) {
	var user models.User
	if err := database.Db.Where("id = ?", c.GetUint("userId")).First(&user).Error; err != nil {
		log.Error(err)
		return
	}
	queue.HandleMessageCreation(&user)
}
