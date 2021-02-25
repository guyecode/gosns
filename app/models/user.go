package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint `json:"id"`
	Nickname string `json:"name"`
	Username string
	Password string
	Age int	`json:"age"`
}

