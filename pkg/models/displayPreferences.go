package models

import "time"

type DisplayPreference struct {
	Id        uint `json:"id"`
	UserId    uint `json:"userid"`
	User      User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DisplayPreferences []DisplayPreference

type DisplayPreferenceInput struct {
	UserId string `json:"userid"`
}
