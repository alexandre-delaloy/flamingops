package models

import "time"

type AwsServicesData struct {
	Id              uint 					`json:"id"`
	UserId    	 		uint 			  `json:"userid"`
	User      	 		User
	Data            string    		`json:"data"`
	CreatedAt       time.Time 		`json:"created_at"`
	UpdatedAt       time.Time 		`json:"updated_at"`
}

type AwsServicesDataInput struct {
	UserId 			uint 		`json:"userid"`
}
