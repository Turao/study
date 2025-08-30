package notification

import "time"

type Model struct {
	Key                 string    `db:"_key"`
	ID                  string    `db:"id"`
	Type                string    `db:"type"`
	Recipient           string    `db:"recipient"`
	Subject             string    `db:"subject"`
	Content             string    `db:"content"`
	CreatedAt           time.Time `db:"created_at"`
	ExternalReferenceID string    `db:"external_reference_id"`
}
