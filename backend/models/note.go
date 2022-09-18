package models

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	Name        string `json:"name"`
	Description string	`json:"description"`
}