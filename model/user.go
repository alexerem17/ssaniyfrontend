package model

import "time"

type User struct {
	ID uint
	Name string `gorm:"not null"`
	Password string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
	Icon string
	CreateTime time.Time `gorm:"DEFAULT:CURRENT_TIMESTAMP"`
}
