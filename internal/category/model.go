package category

type AddCategoryPayload struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type"`
}

type EditCategoryPayload struct {
	ID   uint   `json:"id"`
	Name string `json:"name" binding:"required"`
	Type string `json:"type"`
}

type CategoryPayload struct {
	ID uint `json:"id" binding:"required"`
}

type ListCategoryPayload struct {
	Limit int    `form:"limit" json:"limit"`
	Index int    `form:"index" json:"index"`
	Type  string `form:"type" json:"type"`
}

type CategoryResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
}
