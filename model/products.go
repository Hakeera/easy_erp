package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"not null;size:255"`
	Style       string `gorm:"size:50"` 
	Fabric      string `gorm:"size:50"`
	Color       string `gorm:"size:50"`
	Size        string `gorm:"size:50"`
	Client      string `gorm:"size:50"`  
	Line        string `gorm:"size:50"`
	Category    string `gorm:"size:50"`
	Situation   string `gorm:"size:50"`
	Description string `gorm:"size:500"`
	Price       int64  `gorm:"not null;default:0"`  
	Technical   string `gorm:"size:500"`
}

