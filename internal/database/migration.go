package database

import (
	"github.com/blyndusk/flamingops/pkg/models"
	"github.com/sirupsen/logrus"
)

func Migrate() {
	_ = Db.AutoMigrate(
		&models.User{},
		&models.DisplayPreference{},
		&models.SwServicesData{},
		&models.AwsServicesData{},
		&models.SwCredential{},
		&models.AwsCredential{},
	)
	logrus.Info("Migrations done !")
}
