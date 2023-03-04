package user

import "encoding/json"

type Model struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`

	Tenancy string `json:"tenancy"`

	CreatedAt int64  `json:"created_at"`
	DeletedAt *int64 `json:"deleted_at"`
}

func (m Model) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Model) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}
