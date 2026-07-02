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
	Photo        string    `db:"photo" json:"photo"`
	Instagram    string    `db:"instagram" json:"instagram"`
	Facebook     string    `db:"facebook" json:"facebook"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`

	// Relations
	Category *Category `db:"-" json:"category,omitempty"`
}
