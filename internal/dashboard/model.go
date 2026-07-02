package dashboard

type AddDashboardPayload struct {
	CreatedBy   uint   `db:"created_by" json:"created_by" binding:"required"`
	Title       string `db:"title" json:"title" binding:"required"`
	Description string `db:"description" json:"description" binding:"required"`
	Image       string `db:"image" json:"image" binding:"required"`
}

type EditDashboardPayload struct {
	ID          uint   `db:"id" json:"id" binding:"required"`
	Title       string `db:"title" json:"title" binding:"required"`
	Description string `db:"description" json:"description" binding:"required"`
	Image       string `db:"image" json:"image" binding:"required"`
}

type DashboardPayload struct {
	ID uint `db:"id" json:"id" binding:"required"`
}

type DashboardResponse struct {
	ID          uint   `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Image       string `db:"image" json:"image"`
}
