package message

import (
	"time"

	"github.com/scylladb/gocqlx/v2/table"
)

var _table = table.New(table.Metadata{
	Name: "message",
	Columns: []string{
		"id",
		"author",
		"channel",
		"content",
		"tenancy",
		"created_at",
		"deleted_at",
	},
	PartKey: []string{
		"channel",
	},
	SortKey: []string{
		"created_at",
	},
})

type Model struct {
	ID        string     `json:"id"`
	Author    string     `json:"author"`
	Channel   string     `json:"channel"`
	Content   string     `json:"content"`
	Tenancy   string     `json:"tenancy"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
