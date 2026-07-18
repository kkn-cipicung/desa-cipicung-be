package models

import (
	"time"
)

type Gallery struct {
	ID          uint      `db:"id" json:"id"`
	CreatedBy   uint      `db:"created_by" json:"created_by"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	CategoryID  uint      `db:"category_id" json:"category_id"`
	MediaID     *uint     `db:"media_id" json:"media_id"`
	IsActive    bool      `db:"is_active" json:"is_active"`

	// Relations
	Creator *User  `db:"-" json:"creator,omitempty"`
	Media   *Media `db:"-" json:"media,omitempty"`
}
