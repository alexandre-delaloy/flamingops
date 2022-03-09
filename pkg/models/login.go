package models

type Login struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}
