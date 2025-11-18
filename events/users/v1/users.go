package v1

// UserRegistered is the event for when a user is registered
type UserRegistered struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Tenancy   string `json:"tenancy"`
}
