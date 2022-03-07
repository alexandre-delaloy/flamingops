package models

import "time"

type User struct {
	Id        uint      `json:"id"`
	Username  string    `json:"username"`
	Mail      string      `json:"mail"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User

type UserInput struct {
	Username string `json:"username"`
	Mail     string   `json:"mail"`
	Password string `json:"password"`
}
