package models

import (
	"gorm.io/gorm"
	"gosns/global"
	"time"
)

type Model struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}


func (m Model) Save(){
	global.DB.Save(&m)
}