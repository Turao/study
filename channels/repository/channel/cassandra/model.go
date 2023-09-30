package channel

import (
	"time"

	"github.com/scylladb/gocqlx/v2/table"
)

var _table = table.New(table.Metadata{
	Name: "channel",
	Columns: []string{
		"id",
		"name",
		"tenancy",
		"created_at",
		"deleted_at",
	},
	PartKey: []string{
		"id",
	},
	SortKey: []string{
		"tenancy",
		"created_at",
	},
})

type Model struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Tenancy   string     `json:"tenancy"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
