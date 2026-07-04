package dashboard

import "time"

type AddDashboardPayload struct {
	CreatedBy   uint   `db:"created_by" json:"created_by"`
	CategoryID  uint   `db:"category_id" json:"category_id" binding:"required"`
	Title       string `db:"title" json:"title" binding:"required"`
	Description string `db:"description" json:"description" binding:"required"`
}

type EditDashboardPayload struct {
	ID          uint   `db:"id" json:"id"`
	CategoryID  uint   `db:"category_id" json:"category_id" binding:"required"`
	Title       string `db:"title" json:"title" binding:"required"`
	Description string `db:"description" json:"description" binding:"required"`
}

type DashboardPayload struct {
	ID uint `db:"id" json:"id" binding:"required"`
}

type ListDashboardPayload struct {
	Limit int `form:"limit" json:"limit"`
	Index int `form:"index" json:"index"`
}

type DashboardResponse struct {
	ID          uint      `db:"id" json:"id"`
	CreatedBy   uint      `db:"created_by" json:"created_by"`
	CategoryID  uint      `db:"category_id" json:"category_id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
