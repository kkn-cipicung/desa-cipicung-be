package models

import (
	"time"
)

type Document struct {
	ID            uint      `db:"id" json:"id"`
	CategoryID    uint      `db:"category_id" json:"category_id"`
	UploadedBy    uint      `db:"uploaded_by" json:"uploaded_by"`
	Title         string    `db:"title" json:"title"`
	Description   string    `db:"description" json:"description"`
	File          string    `db:"file" json:"file"`
	DownloadCount int       `db:"download_count" json:"download_count"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`

	// Relations
	Category *Category `db:"-" json:"category,omitempty"`
	Uploader *User     `db:"-" json:"uploader,omitempty"`
}
