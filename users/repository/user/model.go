package user

import (
	"time"
)

// Model is the model for the user entity
type Model struct {
	Key       string     `db:"_key"`
	ID        string     `db:"id"`
	Version   uint32     `db:"version"`
	Email     string     `db:"email"`
	FirstName string     `db:"first_name"`
	LastName  string     `db:"last_name"`
	Tenancy   string     `db:"tenancy"`
	CreatedAt time.Time  `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
