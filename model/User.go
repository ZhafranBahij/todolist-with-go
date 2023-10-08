package model

import "time"

type User struct {
	ID        uint
	Name      string
	Email     string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
