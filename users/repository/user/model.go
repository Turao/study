package user

import (
	"time"
)

type Model struct {
	ID        string     `db:"id"`
	Version   uint32     `db:"version"`
	Email     string     `db:"email"`
	FirstName string     `db:"first_name"`
	LastName  string     `db:"last_name"`
	Tenancy   string     `db:"tenancy"`
	CreatedAt time.Time  `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
