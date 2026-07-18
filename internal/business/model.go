package business

type AddBusinessPayload struct {
	CategoryID   *uint   `json:"category_id" binding:"required"`
	OwnerName    string  `json:"owner_name" binding:"required"`
	BusinessName string  `json:"business_name" binding:"required"`
	Description  string  `json:"description" binding:"required"`
	Phone        string  `json:"phone" binding:"required"`
	Address      string  `json:"address" binding:"required"`
	LocationID   *uint   `json:"location_id"`
	Instagram    *string `json:"instagram"`
	Facebook     *string `json:"facebook"`
}

type EditBusinessPayload struct {
	ID           uint    `json:"id" binding:"required"`
	CategoryID   *uint   `json:"category_id" binding:"required"`
	OwnerName    *string `json:"owner_name"`
	BusinessName *string `json:"business_name"`
	Description  *string `json:"description"`
	Phone        *string `json:"phone"`
	Address      *string `json:"address"`
	LocationID   *uint   `json:"location_id"`
	Instagram    *string `json:"instagram"`
	Facebook     *string `json:"facebook"`
}

type BusinessPayload struct {
	ID uint `json:"id" binding:"required"`
}

type ListBusinessPayload struct {
	Limit int    `form:"limit" json:"limit"`
	Index int    `form:"index" json:"index"`
	Type  string `form:"type" json:"type"`
}

type BusinessResponse struct {
	ID           uint        `json:"id"`
	Category     BusinessRef `json:"category"`
	OwnerName    string      `json:"owner_name"`
	BusinessName string      `json:"business_name"`
	Description  string      `json:"description"`
	Phone        string      `json:"phone"`
	Address      string      `json:"address"`
	LocationID   *uint       `json:"location_id"`
	Instagram    *string     `json:"instagram"`
	Facebook     *string     `json:"facebook"`
	CreatedAt    string      `json:"created_at"`
}

type BusinessRef struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
