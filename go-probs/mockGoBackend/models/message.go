package models

import "time"

type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"index;not null"`
	Body      string    `json:"body" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
}
