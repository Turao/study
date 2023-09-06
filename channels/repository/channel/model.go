package channel

import "time"

type Model struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Tenancy   string     `json:"tenancy"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
