package models

import "time"

type Task struct {
	ID          uint
	Title       string
	Description string
	DueDate     time.Time
	Status      string
}
