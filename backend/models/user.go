package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	
	Name              string    `json:"name"`
	Email 			  string    `json:"email"`
	Password          string    `json:"password"`
	CountryId         uint      `json:"CountryId"`
	Phone             string    `json:"Phone"`
	State             string    `json:"State"`
	IsAdmin           bool      `json:"isAdmin"`
	IsEmailSubscribed bool      `json:"isEmailSubscribed"`
	IsSMSSubscribed   bool      `json:"isSMSSubscribed"`
	
}