package channel

import (
	"time"
)

type Model struct {
	ID        string     `db:"id"`
	Name      string     `db:"name"`
	Tenancy   string     `db:"tenancy"`
	CreatedAt time.Time  `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
