package models

import (
	"time"
)

type ActivityLog struct {
	ID        uint      `db:"id" json:"id"`
	UserID    uint      `db:"user_id" json:"user_id"`
	Activity  string    `db:"activity" json:"activity"`
	IP        string    `db:"ip" json:"ip"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`

	// Relations
	User *User `db:"-" json:"user,omitempty"`
}
