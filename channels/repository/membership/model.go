package membership

import (
	"time"
)

type Model struct {
	Key       string     `db:"_key"`
	ChannelID string     `db:"channel_id"`
	UserID    string     `db:"user_id"`
	Version   uint32     `db:"version"`
	Tenancy   string     `db:"tenancy"`
	CreatedAt time.Time  `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
