package model

type Goal struct {
	ID uint
	UserID uint `gorm:"not null"`
	User User
	Name string `gorm:"not null"`
	Description string
	Mark bool
}

// Get done from all actions
func (g Goal) getProgress() {

}
