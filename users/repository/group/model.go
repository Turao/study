package group

import (
	"time"
)

type Model struct {
	Key       string     `db:"_key"`
	ID        string     `db:"id"`
	Version   uint32     `db:"version"`
	Name      string     `db:"name"`
	Tenancy   string     `db:"tenancy"`
	CreatedAt time.Time  `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
