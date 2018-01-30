package model

import "time"

type Action struct {
	ID uint
	Goal Goal
	Position int
	Name string
	Deadline time.Time
	Checkbox bool
}