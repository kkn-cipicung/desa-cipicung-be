package potential

import "time"

type AddPotentialPayload struct {
	UploadedBy  uint    `db:"uploaded_by" json:"uploaded_by" binding:"required"`
	CategoryID  uint    `db:"category_id" json:"category_id" binding:"required"`
	Title       string  `db:"title" json:"title" binding:"required"`
	Subtitle    string  `db:"subtitle" json:"subtitle"`
	Slug        string  `db:"slug" json:"slug" binding:"required"`
	Description string  `db:"description" json:"description" binding:"required"`
	Latitude    float64 `db:"latitude" json:"latitude" binding:"required"`
	Longitude   float64 `db:"longitude" json:"longitude" binding:"required"`
	OwnerName   string  `db:"owner_name" json:"owner_name" binding:"required"`
	OwnerMsisdn string  `db:"owner_msisdn" json:"owner_msisdn"`
	ImgID       string  `json:"img_id"`
}

type EditPotentialPayload struct {
	ID          uint    `db:"id" json:"id" binding:"required"`
	CategoryID  uint    `db:"category_id" json:"category_id" binding:"required"`
	Title       string  `db:"title" json:"title" binding:"required"`
	Subtitle    string  `db:"subtitle" json:"subtitle"`
	Slug        string  `db:"slug" json:"slug" binding:"required"`
	Description string  `db:"description" json:"description" binding:"required"`
	Latitude    float64 `db:"latitude" json:"latitude" binding:"required"`
	Longitude   float64 `db:"longitude" json:"longitude" binding:"required"`
	OwnerName   string  `db:"owner_name" json:"owner_name" binding:"required"`
	OwnerMsisdn string  `db:"owner_msisdn" json:"owner_msisdn"`
	ImgID       string  `json:"img_id"`
}

type PotentialPayload struct {
	ID uint `db:"id" json:"id" binding:"required"`
}

type ListPotentialPayload struct {
	Limit int `form:"limit" json:"limit"`
	Index int `form:"index" json:"index"`
}

type PotentialResponse struct {
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
}

type potentialMediaPayload struct {
	FilePath   string
	MimeType   string
	UploadedBy *uint
}
