package model

import "time"

type Action struct {
	ID uint
	GoalID uint
	Weight int
	Name string
	Deadline time.Time
	Done bool
	CRON string
}

// Get mark from Goal
func (a Action) mark() {

}