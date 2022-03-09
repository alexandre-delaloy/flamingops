package models

import "time"

type RequestedRegions struct {
	Id        uint `json:"id"`
	UserId    uint `json:"userid"`
	User      User
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RequestedRegionsInput struct {
	UserId  uint   `json:"userid"`
	Content string `json:"content"`
}
