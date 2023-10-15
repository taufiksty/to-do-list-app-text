package entity

import "time"

type Task struct {
	Id                 int32
	Title, Description string
	Done               bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
