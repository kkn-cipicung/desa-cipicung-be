package models

import (
	"time"
)

type Category struct {
	ID        uint      `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Slug      string    `db:"slug" json:"slug"` // url friendly
	Type      string    `db:"type" json:"type"` // news, document, business
	CreatedAt time.Time `db:"created_at" json:"created_at"`

	// Relations
	Posts      []Post      `db:"-" json:"posts,omitempty"`
	Documents  []Document  `db:"-" json:"documents,omitempty"`
	Businesses []Business  `db:"-" json:"businesses,omitempty"`
	Potentials []Potential `db:"-" json:"potentials,omitempty"`
}
