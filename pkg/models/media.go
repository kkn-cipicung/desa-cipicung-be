package models

import (
	"time"
)

type Media struct {
	ID           uint      `db:"id" json:"id"`
	UploadedBy   uint      `db:"uploaded_by" json:"uploaded_by"`
	EntityType   string    `db:"entity_type" json:"entity_type"`
	EntityID     uint      `db:"entity_id" json:"entity_id"`
	Role         string    `db:"role" json:"role"`
	OriginalName string    `db:"original_name" json:"original_name"`
	FileName     string    `db:"file_name" json:"file_name"`
	FilePath     string    `db:"file_path" json:"file_path"`
	MimeType     string    `db:"mime_type" json:"mime_type"`
	FileSize     int64     `db:"file_size" json:"file_size"`
	Caption      string    `db:"caption" json:"caption"`
	OrderNumber  int       `db:"order_number" json:"order_number"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`

	// Relations
	Uploader *User `db:"-" json:"uploader,omitempty"`
}
