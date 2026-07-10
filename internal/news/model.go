package news

import "time"

type AddNewsPayload struct {
	UploadedBy  uint   `db:"uploaded_by" json:"uploaded_by"`
	CategoryID  uint   `db:"category_id" json:"category_id" binding:"required"`
	Title       string `db:"title" json:"title" binding:"required"`
	Description string `db:"description" json:"description" binding:"required"`
	ImgID       string `json:"img_id"`
}

type EditNewsPayload struct {
	ID          uint   `db:"id" json:"id" binding:"required"`
	CategoryID  uint   `db:"category_id" json:"category_id" binding:"required"`
	Title       string `db:"title" json:"title" binding:"required"`
	Description string `db:"description" json:"description" binding:"required"`
	ImgID       string `json:"img_id"`
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
	ID          uint      `db:"id" json:"id"`
	CategoryID  uint      `db:"category_id" json:"category_id"`
	UploadedBy  uint      `db:"uploaded_by" json:"uploaded_by"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
