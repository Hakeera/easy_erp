package model

import "gorm.io/gorm"

type Cliente struct {
	gorm.Model
	Nome  string `gorm:"type:varchar(255);not null"`
	Email string `gorm:"type:varchar(255)"`
	Telefone string `gorm:"type:varchar(50)"`
}

