package metadata

import "time"

// Auditable is the interface for auditable entities
type Auditable interface {
	CreatedAt() time.Time
}

// Deletable is the interface for deletable entities
type Deletable interface {
	Delete()
	DeletedAt() *time.Time
}
