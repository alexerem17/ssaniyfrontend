package model

type User struct {
	ID uint
	Name string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Icon string
}
