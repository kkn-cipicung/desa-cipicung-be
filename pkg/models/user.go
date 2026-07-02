package models

import (
	"time"
)

type User struct {
	ID        uint       `db:"id" json:"id"`
	Role      string     `db:"role" json:"role"` // superadmin, admin
	Name      string     `db:"name" json:"name"`
	Username  string     `db:"username" json:"username"`
	Password  string     `db:"password" json:"password"`
	IsActive  bool       `db:"is_active" json:"is_active"`
	LastLogin *time.Time `db:"last_login" json:"last_login"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`

	// Relations
	Posts        []Post        `db:"-" json:"posts,omitempty"`
	Events       []Event       `db:"-" json:"events,omitempty"`
	Galleries    []Gallery     `db:"-" json:"galleries,omitempty"`
	Documents    []Document    `db:"-" json:"documents,omitempty"`
	ActivityLogs []ActivityLog `db:"-" json:"activity_logs,omitempty"`
}
