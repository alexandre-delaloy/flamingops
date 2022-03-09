package forms

import (
	"basic-api/models"
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)



// LoginForm is used to log the user in
type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
