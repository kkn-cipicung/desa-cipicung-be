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
	CategoryID  uint      `db:"category_id" json:"category_id"`

	// Relations
	Creator *User   `db:"-" json:"creator,omitempty"`
	Media   []Media `db:"-" json:"media,omitempty"`
}
