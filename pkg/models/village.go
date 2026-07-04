package models

import (
	"time"
)

type Village struct {
	ID          uint      `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Province    string    `db:"province" json:"province"`
	Regency     string    `db:"regency" json:"regency"`
	District    string    `db:"district" json:"district"`
	PostalCode  string    `db:"postal_code" json:"postal_code"`
	Address     string    `db:"address" json:"address"`
	Phone       string    `db:"phone" json:"phone"`
	Email       string    `db:"email" json:"email"`
	Website     string    `db:"website" json:"website"`
	Latitude    float64   `db:"latitude" json:"latitude"`
	Longitude   float64   `db:"longitude" json:"longitude"`
	Vision      string    `db:"vision" json:"vision"`
	Mission     string    `db:"mission" json:"mission"`
	History     string    `db:"history" json:"history"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`

	// Relations
	Officials []Official `db:"-" json:"officials,omitempty"`
	Media     []Media    `db:"-" json:"media,omitempty"`
}
