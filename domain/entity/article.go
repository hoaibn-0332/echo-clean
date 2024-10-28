package entity

import (
	"time"
)

// Article is representing the Article data struct
type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" conform:"trim" validate:"required"`
	Content   string    `json:"content" conform:"trim" validate:"required"`
	Author    Author    `json:"author"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
