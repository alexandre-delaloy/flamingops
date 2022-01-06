package database

import (
	"github.com/blyndusk/flamingops/pkg/models"
	"github.com/sirupsen/logrus"
)

func Migrate() {
	_ = Db.AutoMigrate(&models.User{})
	logrus.Info("Migrations done !")
}
