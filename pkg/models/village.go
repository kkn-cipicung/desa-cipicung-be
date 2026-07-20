package models

import (
	"time"

	"github.com/lib/pq"
)

type Village struct {
	ID          uint           `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Province    string         `db:"province" json:"province"`
	Regency     string         `db:"regency" json:"regency"`
	District    string         `db:"district" json:"district"`
	PostalCode  string         `db:"postal_code" json:"postal_code"`
	Address     string         `db:"address" json:"address"`
	Phone       string         `db:"phone" json:"phone"`
	Email       string         `db:"email" json:"email"`
	Latitude    float64        `db:"latitude" json:"latitude"`
	Longitude   float64        `db:"longitude" json:"longitude"`
	Vision      string         `db:"vision" json:"vision"`
	Mission     pq.StringArray `db:"mission" json:"mission"`
	History     string         `db:"history" json:"history"`
	Description string         `db:"description" json:"description"`
	Region      string         `db:"region" json:"region"`
	HamletOne   string         `db:"hamlet_one" json:"hamlet_one"`
	HamletTwo   string         `db:"hamlet_two" json:"hamlet_two"`
	NorthBorder string         `db:"north_border" json:"north_border"`
	EastBorder  string         `db:"east_border" json:"east_border"`
	SouthBorder string         `db:"south_border" json:"south_border"`
	WestBorder  string         `db:"west_border" json:"west_border"`
	Area        string         `db:"area" json:"area"`
	Elevation   string         `db:"elevation" json:"elevation"`
	Coordinate  string         `db:"coordinate" json:"coordinate"`
	Population  string         `db:"population" json:"population"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`

	// Relations
	Officials []Official `db:"-" json:"officials,omitempty"`
	Media     []Media    `db:"-" json:"media,omitempty"`
}
