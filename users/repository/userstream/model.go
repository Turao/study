package user

// Event is a CDC event for users table
type Event struct {
	Payload EventPayload `json:"payload"`
}

// EventPayload is the paylod of the CDC event
type EventPayload struct {
	Before *Model `json:"before"`
	After  *Model `json:"after"`
}

// Model is the model for the user entity
type Model struct {
	Key       string `json:"_key"`
	ID        string `json:"id"`
	Version   uint32 `json:"version"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Tenancy   string `json:"tenancy"`
	CreatedAt int64  `json:"created_at"`
	DeletedAt *int64 `json:"deleted_at"`
}
