package model

import "time"

type Action struct {
	ID uint
	Name string `gorm:"not null"`
	Done bool `gorm:"not null;DEFAULT:0"`
	Rating uint
	Repeater string
	Deadline time.Time
	GPS gps `gorm:"type:point"`

	CreateTime time.Time `gorm:"DEFAULT:CURRENT_TIMESTAMP"`
	// Foreign keys
	GoalID uint
	Goal Goal
	UserID uint
	User Goal
}

type gps struct {
	x float64
	y float64
}