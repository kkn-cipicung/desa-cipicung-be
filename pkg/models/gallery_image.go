package models

type GalleryImage struct {
	ID          uint   `db:"id" json:"id"`
	GalleryID   uint   `db:"gallery_id" json:"gallery_id"`
	Image       string `db:"image" json:"image"`
	Caption     string `db:"caption" json:"caption"`
	OrderNumber int    `db:"order_number" json:"order_number"`

	// Relations
	Gallery *Gallery `db:"-" json:"gallery,omitempty"`
}
