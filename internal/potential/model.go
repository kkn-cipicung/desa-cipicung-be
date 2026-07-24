package potential

import "time"

type AddPotentialPayload struct {
	UploadedBy  uint                    `db:"uploaded_by" json:"uploaded_by"`
	CategoryID  uint                    `db:"category_id" json:"category_id" binding:"required"`
	Title       string                  `db:"title" json:"title" binding:"required"`
	Subtitle    string                  `db:"subtitle" json:"subtitle"`
	Slug        string                  `db:"slug" json:"slug" binding:"required"`
	Description string                  `db:"description" json:"description" binding:"required"`
	LocationID  *uint                   `db:"location_id" json:"location_id"`
	Location    *PotentialLocationInput `json:"location"`
	// OwnerName   string                  `db:"owner_name" json:"owner_name" binding:"required"`
	// OwnerMsisdn string                  `db:"owner_msisdn" json:"owner_msisdn"`
	MediaID *string `json:"media_id"`
}

type EditPotentialPayload struct {
	ID          uint                    `db:"id" json:"id" binding:"required"`
	UploadedBy  uint                    `db:"-" json:"-"`
	CategoryID  uint                    `db:"category_id" json:"category_id" binding:"required"`
	Title       string                  `db:"title" json:"title" binding:"required"`
	Subtitle    string                  `db:"subtitle" json:"subtitle"`
	Slug        string                  `db:"slug" json:"slug" binding:"required"`
	Description string                  `db:"description" json:"description" binding:"required"`
	LocationID  *uint                   `db:"location_id" json:"location_id"`
	Location    *PotentialLocationInput `json:"location"`
	// OwnerName   string                  `db:"owner_name" json:"owner_name" binding:"required"`
	// OwnerMsisdn string                  `db:"owner_msisdn" json:"owner_msisdn"`
	MediaID *string `json:"media_id"`
}

type PotentialPayload struct {
	ID uint `db:"id" json:"id" binding:"required"`
}

type PotentialLocationInput struct {
	ID          *uint   `json:"id"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
}

type ListPotentialPayload struct {
	Limit int `form:"limit" json:"limit"`
	Index int `form:"index" json:"index"`
}

type PotentialResponse struct {
	ID           uint   `db:"id" json:"id"`
	CategoryID   uint   `db:"category_id" json:"category_id"`
	CategoryName string `db:"category_name" json:"category_name"`
	Title        string `db:"title" json:"title"`
	Subtitle     string `db:"subtitle" json:"subtitle"`
	Slug         string `db:"slug" json:"slug"`
	Description  string `db:"description" json:"description"`
	LocationID   *uint  `db:"location_id" json:"location_id"`
	// OwnerName    string     `db:"owner_name" json:"owner_name"`
	// OwnerMsisdn  string     `db:"owner_msisdn" json:"owner_msisdn"`
	Media     string     `db:"media" json:"media"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
}

type PotentialOutput struct {
	ID          uint               `json:"id"`
	Category    PotentialRef       `json:"category"`
	Title       string             `json:"title"`
	Subtitle    string             `json:"subtitle"`
	Slug        string             `json:"slug"`
	Description string             `json:"description"`
	Location    *PotentialLocation `json:"location"`
	Owner       PotentialOwner     `json:"owner"`
	Media       string             `json:"media"`
	CreatedAt   string             `json:"created_at"`
}

type PotentialRef struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PotentialLocation struct {
	ID uint `json:"id"`
}

type PotentialOwner struct {
	Name   string `json:"name"`
	Msisdn string `json:"msisdn"`
}
