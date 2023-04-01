package message

import "time"

type Model struct {
	ID        string     `json:"id"`
	Channel   string     `json:"channel"`
	Content   string     `json:"content"`
	Tenancy   string     `json:"tenancy"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
