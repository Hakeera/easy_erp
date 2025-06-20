package model

import "time"


type User struct {
    ID       uint64 `gorm:"primaryKey;autoIncrement"`
    Username string `gorm:"not null;unique;size:100"`
    Email    string `gorm:"not null;unique;size:255"`
    Password string `gorm:"not null;size:255"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
