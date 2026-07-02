package models

import (
	"time"
)

type Gallery struct {
	ID          uint      `db:"id" json:"id"`
	CreatedBy   uint      `db:"created_by" json:"created_by"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Cover       string    `db:"cover" json:"cover"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`

	// Relations
	Author *User          `db:"-" json:"author,omitempty"`
	Images []GalleryImage `db:"-" json:"images,omitempty"`
}
