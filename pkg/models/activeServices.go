package models

import "time"

type ActiveServices struct {
	Id          uint `json:"id"`
	UserId      uint `json:"userid"`
	User        User
	AwsServices string    `json:"awsservices"`
	SwServices  string    `json:"swservices"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ActiveServicesInput struct {
	UserId      uint   `json:"userid"`
	AwsServices string `json:"awsservices"`
	SwServices  string `json:"swservices"`
}
