package gallery

type AddGalleryPayload struct {
	CreatedBy   uint    `db:"created_by" json:"created_by"`
	CategoryID  uint    `db:"category_id" json:"category_id" binding:"required"`
	Title       string  `db:"title" json:"title" binding:"required"`
	Description string  `db:"description" json:"description" binding:"required"`
	MediaID     *string `json:"media_id"`
}

type EditGalleryPayload struct {
	ID          uint    `db:"id" json:"id" binding:"required"`
	UpdatedBy   uint    `db:"-" json:"-"`
	CategoryID  uint    `db:"category_id" json:"category_id" binding:"required"`
	Title       string  `db:"title" json:"title" binding:"required"`
	Description string  `db:"description" json:"description" binding:"required"`
	MediaID     *string `json:"media_id"`
}

type GalleryPayload struct {
	ID uint `db:"id" json:"id" binding:"required"`
}

type ListGalleryPayload struct {
	Limit int `form:"limit" json:"limit"`
	Index int `form:"index" json:"index"`
}

type GalleryResponse struct {
	ID           uint   `db:"id" json:"id"`
	CategoryID   uint   `db:"category_id" json:"category_id"`
	CategoryName string `db:"category_name" json:"category_name"`
	Title        string `db:"title" json:"title"`
	Description  string `db:"description" json:"description"`
	Image        string `db:"image" json:"image"`
}

type GalleryListOutput struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}

type GalleryDetailOutput struct {
	Title       string       `json:"title"`
	Image       string       `json:"image"`
	Description string       `json:"description"`
	Category    []GalleryRef `json:"category"`
}

type GalleryRef struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
