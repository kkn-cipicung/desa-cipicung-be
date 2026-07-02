package models

import (
	"time"
)

type Event struct {
	ID          uint       `db:"id" json:"id"`
	AuthorID    uint       `db:"author_id" json:"author_id"`
	Title       string     `db:"title" json:"title"`
	Slug        string     `db:"slug" json:"slug"`
	Description string     `db:"description" json:"description"`
	Location    string     `db:"location" json:"location"`
	StartDate   *time.Time `db:"start_date" json:"start_date"`
	EndDate     *time.Time `db:"end_date" json:"end_date"`
	Thumbnail   string     `db:"thumbnail" json:"thumbnail"`
	Status      string     `db:"status" json:"status"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`

	// Relations
	Author *User        `db:"-" json:"author,omitempty"`
	Images []EventImage `db:"-" json:"images,omitempty"`
}
