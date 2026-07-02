package models

type EventImage struct {
	ID          uint   `db:"id" json:"id"`
	EventID     uint   `db:"event_id" json:"event_id"`
	Image       string `db:"image" json:"image"`
	Caption     string `db:"caption" json:"caption"`
	OrderNumber int    `db:"order_number" json:"order_number"`

	// Relations
	Event *Event `db:"-" json:"event,omitempty"`
}
