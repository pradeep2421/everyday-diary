package models

import "time"

type User struct {
	Id                uint      `json:"id" gorm:"primary_key;auto_increment"`
	Name              string    `json:"name"`
	Password          string    `json:"password"`
	CountryId         uint      `json:"CountryId"`
	Phone             string    `json:"Phone"`
	State             string    `json:"State"`
	IsAdmin           bool      `json:"isAdmin"`
	IsEmailSubscribed bool      `json:"isEmailSubscribed"`
	IsSMSSubscribed   bool      `json:"isSMSSubscribed"`
	CreatedAt         time.Time `json:"CreatedAt"`
}