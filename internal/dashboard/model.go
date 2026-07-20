package dashboard

import "time"

type AddDashboardPayload struct {
	CreatedBy   uint    `db:"created_by" json:"created_by"`
	CategoryID  uint    `db:"category_id" json:"category_id" binding:"required"`
	Title       string  `db:"title" json:"title" binding:"required"`
	Description string  `db:"description" json:"description" binding:"required"`
	MediaID     *string `json:"media_id"`
}

type EditDashboardPayload struct {
	ID          uint    `db:"id" json:"id"`
	UpdatedBy   uint    `db:"-" json:"-"`
	CategoryID  uint    `db:"category_id" json:"category_id" binding:"required"`
	Title       string  `db:"title" json:"title" binding:"required"`
	Description string  `db:"description" json:"description" binding:"required"`
	MediaID     *string `json:"media_id"`
	IsActive    bool    `db:"is_active" json:"is_active"`
}

type DashboardPayload struct {
	ID uint `db:"id" json:"id" binding:"required"`
}

type ListDashboardPayload struct {
	Limit int `form:"limit" json:"limit"`
	Index int `form:"index" json:"index"`
}

type DashboardResponse struct {
	ID           uint      `db:"id" json:"id"`
	CreatedBy    uint      `db:"created_by" json:"created_by"`
	CreatorName  string    `db:"creator_name" json:"creator_name"`
	CategoryID   uint      `db:"category_id" json:"category_id"`
	CategoryName string    `db:"category_name" json:"category_name"`
	Title        string    `db:"title" json:"title"`
	Description  string    `db:"description" json:"description"`
	MediaID      *uint     `db:"media_id" json:"media_id"`
	IsActive     bool      `db:"is_active" json:"is_active"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

type DashboardOutput struct {
	ID          uint         `json:"id"`
	Creator     DashboardRef `json:"creator"`
	Category    DashboardRef `json:"category"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	MediaID     *uint        `json:"media_id"`
	IsActive    bool         `json:"is_active"`
	CreatedAt   string       `json:"created_at"`
}

type DashboardRef struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
