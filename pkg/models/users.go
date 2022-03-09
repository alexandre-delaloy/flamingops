package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint      `json:"id"`
	Username  string    `json:"username"`
	Mail      string    `json:"mail"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User

type UserInput struct {
	Username string `json:"username"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

func (u *User) PrepareResponse() {
	u.Password = ""
}

func (u User) CheckPassword(pswrd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pswrd))
	return err
}
