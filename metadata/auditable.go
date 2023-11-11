package metadata

import "time"

type Auditable interface {
	CreatedAt() time.Time
}

type Deletable interface {
	Delete()
	DeletedAt() *time.Time
}
