package models

import (
	"time"
)

type Post struct {
	ID           uint       `db:"id" json:"id"`
	CategoryID   uint       `db:"category_id" json:"category_id"`
	AuthorID     uint       `db:"author_id" json:"author_id"`
	Type         string     `db:"type" json:"type"` // news, announcement, potential, article
	Title        string     `db:"title" json:"title"`
	Slug         string     `db:"slug" json:"slug"`
	Excerpt      string     `db:"excerpt" json:"excerpt"`
	Content      string     `db:"content" json:"content"`
	PublishStart *time.Time `db:"publish_start" json:"publish_start"`
	PublishEnd   *time.Time `db:"publish_end" json:"publish_end"`
	IsPinned     bool       `db:"is_pinned" json:"is_pinned"`
	Status       string     `db:"status" json:"status"`
	Views        int        `db:"views" json:"views"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`

	// Relations
	Category *Category `db:"-" json:"category,omitempty"`
	Author   *User     `db:"-" json:"author,omitempty"`
	Media    []Media   `db:"-" json:"media,omitempty"`
}
