package model

type Action struct {
	ID uint
	GoalID uint
	Goal Goal
	Name string `gorm:"not null"`
	Done bool `gorm:"not null;DEFAULT:0"`
	CRON string
	GPS gps `gorm:"type:point"`
}

type gps struct {
	x float64
	y float64
}