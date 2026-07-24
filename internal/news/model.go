package news

import "time"

type AddNewsPayload struct {
	UploadedBy  uint    `db:"uploaded_by" json:"uploaded_by"`
	CategoryID  uint    `db:"category_id" json:"category_id" binding:"required"`
	Title       string  `db:"title" json:"title" binding:"required"`
	Description string  `db:"description" json:"description" binding:"required"`
	MediaID     *string `json:"media_id"`
}

type EditNewsPayload struct {
	ID          uint    `db:"id" json:"id" binding:"required"`
	UploadedBy  uint    `db:"-" json:"-"`
	CategoryID  uint    `db:"category_id" json:"category_id" binding:"required"`
	Title       string  `db:"title" json:"title" binding:"required"`
	Description string  `db:"description" json:"description" binding:"required"`
	MediaID     *string `json:"media_id"`
}

type NewsPayload struct {
	ID uint `db:"id" json:"id" binding:"required"`
}

type ListNewsPayload struct {
	Limit int `form:"limit" json:"limit"`
	Index int `form:"index" json:"index"`
}

type NewsByIdPayload struct {
	ID uint `db:"id" json:"id" binding:"required"`
}

type NewsByDatePayload struct {
	Date string `json:"date" binding:"required"`
}

type NewsResponse struct {
	ID           uint      `db:"id" json:"id"`
	CategoryID   uint      `db:"category_id" json:"category_id"`
	CategoryName string    `db:"category_name" json:"category_name"`
	UploadedBy   uint      `db:"uploaded_by" json:"uploaded_by"`
	UploaderName string    `db:"uploader_name" json:"uploader_name"`
	Title        string    `db:"title" json:"title"`
	Description  string    `db:"description" json:"description"`
	Media        string    `db:"media" json:"media"`
	Source       string    `db:"source" json:"source"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

type NewsHeaderResponse struct {
	ID    uint   `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}

type NewsOutput struct {
	ID           uint    `json:"id"`
	CategoryID   uint    `json:"category_id"`
	CategoryName string  `json:"category_name"`
	Category     NewsRef `json:"category"`
	Uploader     NewsRef `json:"uploader"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Media        string  `json:"media"`
	Source       string  `json:"source"`
	CreatedAt    string  `json:"created_at"`
}

type NewsRef struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
