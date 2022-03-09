package models

import "time"

type SwServicesData struct {
	Id        uint `json:"id"`
	UserId    uint `json:"userid"`
	User      User
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SwServicesDataInput struct {
	UserId uint `json:"userid"`
}
