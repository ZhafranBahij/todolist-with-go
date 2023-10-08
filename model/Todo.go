package model

import "time"

type Todo struct {
	ID        uint
	UserID    uint
	User      User `gorm:"foreignKey:UserID"`
	Event     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
