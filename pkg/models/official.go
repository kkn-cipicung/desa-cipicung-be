package models

import (
	"time"
)

type Official struct {
	ID          uint      `db:"id" json:"id"`
	VillageID   uint      `db:"village_id" json:"village_id"`
	Name        string    `db:"name" json:"name"`
	Position    string    `db:"position" json:"position"`
	Phone       string    `db:"phone" json:"phone"`
	Email       string    `db:"email" json:"email"`
	Description string    `db:"description" json:"description"`
	OrderNumber int       `db:"order_number" json:"order_number"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`

	// Relations
	Village *Village `db:"-" json:"village,omitempty"`
	Media   []Media  `db:"-" json:"media,omitempty"`
}
