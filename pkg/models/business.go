package models

import (
	"time"
)

type Business struct {
	ID           uint      `db:"id" json:"id"`
	CategoryID   uint      `db:"category_id" json:"category_id"`
	OwnerName    string    `db:"owner_name" json:"owner_name"`
	BusinessName string    `db:"business_name" json:"business_name"`
	Description  string    `db:"description" json:"description"`
	Phone        string    `db:"phone" json:"phone"`
	Address      string    `db:"address" json:"address"`
	LocationID   *uint     `db:"location_id" json:"location_id"`
	MediaID      *uint     `db:"media_id" json:"media_id"`
	Instagram    *string   `db:"instagram" json:"instagram"`
	Facebook     *string   `db:"facebook" json:"facebook"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	CategoryName string    `db:"category_name" json:"category_name"`

	// Relations
	Location *Location `db:"-" json:"location,omitempty"`
	Category *Category `db:"-" json:"category,omitempty"`
	Media    *Media    `db:"-" json:"media,omitempty"`
}
