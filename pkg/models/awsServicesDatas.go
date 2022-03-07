package models

import "time"

type AwsServicesData struct {
	Id              uint 					`json:"id"`
	AwsCredentialId uint 					`json:"awscredentialid"`
	AwsCredential   AwsCredential
	Data            string    		`json:"data"`
	CreatedAt       time.Time 		`json:"created_at"`
	UpdatedAt       time.Time 		`json:"updated_at"`
}

type AwsServicesDatas []AwsServicesData

type AwsServicesDataInput struct {
	AwsCredentialId string `json:"awscredentialid"`
	Data            string `json:"data"`
}
