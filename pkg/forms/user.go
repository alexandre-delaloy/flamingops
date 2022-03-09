package forms

import (
	"basic-api/models"
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// UserForm is used to edit or create an user
type UserForm struct {
	Dob             time.Time `json:"dob"`
	Email           string    `json:"email"`
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirm_password"`
}

// LoginForm is used to log the user in
type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u UserForm) PrepareUser() (*models.User, error) {
	if u.Password != u.ConfirmPassword {
		log.Println("passwords don't match")
		return nil, errors.New("passwords don't match")
	}

	hashedPassword, err := (bcrypt.GenerateFromPassword([]byte(u.Password), 10))
	if err != nil {
		return nil, err
	}

	dob, err := time.Parse("2006-01-02", u.Dob.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	user := models.User{
		Username: u.Username,
		Password: string(hashedPassword),
		Dob:      dob,
		Email:    u.Email,
	}
	return &user, nil
}
