package notification

import "time"

type Notification struct {
	ID        string                 `json:"id"`
	Type      string                 `json:"type"`
	Recipient string                 `json:"recipient"`
	Subject   string                 `json:"subject"`
	Content   map[string]interface{} `json:"content"`
	CreatedAt time.Time              `json:"created_at"`

	ExternalReferenceID string `json:"external_reference_id"`
}

type Type string

const (
	TypeConfirmation Type = "confirmation"
)
