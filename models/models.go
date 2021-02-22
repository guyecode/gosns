package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint `json:"id"`
	Name string `json:"name"`
	Age int	`json:"age"`
}

