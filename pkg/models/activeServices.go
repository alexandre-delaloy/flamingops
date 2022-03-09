package models

import "time"

type ActiveServices struct {
	Id          uint `json:"id"`
	UserId      uint `json:"userid"`
	User        User
	AwsServices string    `json:"content"`
	SwServices  string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ActiveServicesInput struct {
	UserId      uint   `json:"userid"`
	AwsServices string `json:"content"`
	SwServices  string `json:"content"`
}
