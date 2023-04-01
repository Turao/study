package message

import "time"

type Model struct {
	ID        string     `json:"id"`
	Content   string     `json:"content"`
	Channels  []string   `json:"channels"`
	Tenancy   string     `json:"tenancy"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
