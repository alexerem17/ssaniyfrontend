package model

import "time"

type Goal struct {
	ID uint
	Name string `gorm:"not null"`
	Description string
	Mark bool
	CreateTime time.Time `gorm:"DEFAULT:CURRENT_TIMESTAMP"`

	// Foreign key
	UserID uint
	User User
}

// Get done from all actions
func (g Goal) getProgress() {

}
