package models

import "time"

// User is a GORM model. Struct tags control JSON serialization and DB column mapping.
//
// Go's static typing means every field has a known type — no runtime type surprises
// when this crosses handler → service → repository boundaries.
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
}
