package models

import "time"

type SwServicesData struct {
	Id             uint `json:"id"`
	SwCredentialId uint `json:"swcredentialid"`
	SwCredential   SwCredential
	Data           string    `json:"data"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SwServicesDatas []SwServicesData

type SwServicesDataInput struct {
	SwCredentialId string `json:"swcredentialid"`
}
