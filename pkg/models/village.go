package models

import (
	"time"

	"cipicung.id/be/pkg/types"
)

type Village struct {
	ID          uint                  `db:"id" json:"id"`
	Name        string                `db:"name" json:"name"`
	Province    string                `db:"province" json:"province"`
	Regency     string                `db:"regency" json:"regency"`
	District    string                `db:"district" json:"district"`
	PostalCode  string                `db:"postal_code" json:"postal_code"`
	Address     string                `db:"address" json:"address"`
	Phone       string                `db:"phone" json:"phone"`
	Email       string                `db:"email" json:"email"`
	Instagram   string                `db:"ig_usn" json:"ig_usn"`
	TikTok      string                `db:"tiktok_usn" json:"tiktok_usn"`
	YouTube     string                `db:"yt_usn" json:"yt_usn"`
	Latitude    float64               `db:"latitude" json:"latitude"`
	Longitude   float64               `db:"longitude" json:"longitude"`
	Vision      string                `db:"vision" json:"vision"`
	Mission     types.JSONStringArray `db:"mission" json:"mission"`
	History     string                `db:"history" json:"history"`
	Description string                `db:"description" json:"description"`
	Region      string                `db:"region" json:"region"`
	HamletOne   int64                 `db:"hamlet_one" json:"hamlet_one"`
	HamletTwo   int64                 `db:"hamlet_two" json:"hamlet_two"`
	TotalRT     int64                 `db:"total_rt" json:"total_rt"`
	TotalRW     int64                 `db:"total_rw" json:"total_rw"`
	RTHamletOne int64                 `db:"rt_hamlet_one" json:"rt_hamlet_one"`
	RTHamletTwo int64                 `db:"rt_hamlet_two" json:"rt_hamlet_two"`
	RWHamletOne int64                 `db:"rw_hamlet_one" json:"rw_hamlet_one"`
	RWHamletTwo int64                 `db:"rw_hamlet_two" json:"rw_hamlet_two"`
	NorthBorder string                `db:"north_border" json:"north_border"`
	EastBorder  string                `db:"east_border" json:"east_border"`
	SouthBorder string                `db:"south_border" json:"south_border"`
	WestBorder  string                `db:"west_border" json:"west_border"`
	Area        string                `db:"area" json:"area"`
	Elevation   string                `db:"elevation" json:"elevation"`
	Coordinate  string                `db:"coordinate" json:"coordinate"`
	Population  string                `db:"population" json:"population"`
	TotalMale   int64                 `db:"total_male" json:"total_male"`
	TotalFemale int64                 `db:"total_female" json:"total_female"`
	CreatedAt   time.Time             `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time             `db:"updated_at" json:"updated_at"`

	// Relations
	Officials []Official `db:"-" json:"officials,omitempty"`
	Media     []Media    `db:"-" json:"media,omitempty"`
}
