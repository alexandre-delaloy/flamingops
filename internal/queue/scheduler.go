package queue

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/blyndusk/flamingops/internal/database"
	"github.com/blyndusk/flamingops/pkg/models"
)

func main() {
	ticker := time.NewTicker(24 * time.Hour)
	for ; true; <-ticker.C {
		sendMessages()
	}
}

func sendMessages() {
	var users []models.User
	if err := database.Db.Find(&users).Error; err != nil {
		log.Error(err)
		return
	}
	for _, user := range users {
		HandleMessageCreation(&user)
	}
}
