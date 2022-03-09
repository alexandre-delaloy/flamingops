package models

import "time"

type Login struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}
