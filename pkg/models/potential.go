package models

import (
	"time"
)

type Potential struct {
	ID          uint      `db:"id" json:"id"`
	CategoryID  uint      `db:"category_id" json:"category_id"`
	Title       string    `db:"title" json:"title"`
	Subtitle    string    `db:"subtitle" json:"subtitle"`
	Slug        string    `db:"slug" json:"slug"`
	Description string    `db:"description" json:"description"`
	LocationID  uint      `db:"location_id" json:"location_id"`
	OwnerName   string    `db:"owner_name" json:"owner_name"`
	OwnerMsisdn string    `db:"owner_msisdn" json:"owner_msisdn"`
	MediaId     uint      `db:"media_id" json:"media_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`

	Category *Category `db:"-" json:"category,omitempty"`
	Media    []Media   `db:"-" json:"media,omitempty"`
}
