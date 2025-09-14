package notification

import "time"

type Model struct {
	Key       string    `db:"_key"`
	ID        string    `db:"id"`
	Type      string    `db:"type"`
	Recipient string    `db:"recipient"`
	Subject   string    `db:"subject"`
	Content   string    `db:"content"`
	Metadata  string    `db:"metadata"`
	CreatedAt time.Time `db:"created_at"`
}
