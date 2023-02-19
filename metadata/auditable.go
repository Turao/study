package metadata

import "time"

type Auditable interface {
	CreatedAt() time.Time
	DeletedAt() *time.Time
}
