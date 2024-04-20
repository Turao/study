package group

import (
	"time"
)

type GroupModel struct {
	Key       string     `db:"_key"`
	ID        string     `db:"id"`
	Version   uint32     `db:"version"`
	Name      string     `db:"name"`
	Tenancy   string     `db:"tenancy"`
	CreatedAt time.Time  `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type GroupMemberModel struct {
	Key          string `db:"_key"`
	GroupID      string `db:"group_id"`
	GroupVersion uint32 `db:"group_version"`
	MemberID     string `db:"member_id"`
}
