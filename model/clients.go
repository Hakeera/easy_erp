package model

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255);not null"`
	Email string `gorm:"type:varchar(255)"`
	Phone string `gorm:"type:varchar(50)"`
}

