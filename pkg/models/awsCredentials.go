package models

import "time"

type AwsCredential struct {
	Id          uint 			`json:"id"`
	UserId      uint 			`json:"userid"`
	User        User
	Credentials string    `json:"credentials"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type AwsCredentials []AwsCredential

type AwsCredentialInput struct {
	UserId      string `json:"userid"`
	Credentials string `json:"credentials"`
}
