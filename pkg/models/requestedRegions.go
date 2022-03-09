package models

import "time"

type RequestedRegions struct {
	Id        uint `json:"id"`
	UserId    uint `json:"userid"`
	User      User
	AwsRegion string    `json:"awsregion"`
	SwRegion  string    `json:"swregion"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RequestedRegionsInput struct {
	UserId    uint   `json:"userid"`
	AwsRegion string `json:"awsregion"`
	SwRegion  string `json:"swregion"`
}
