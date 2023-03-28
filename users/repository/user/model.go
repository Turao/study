package user

import (
	"encoding/json"
	"time"
)

type Model struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	Tenancy string `json:"tenancy"`

	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (m Model) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Model) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}
