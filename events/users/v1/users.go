package v1

type UserRegistered struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Tenancy   string `json:"tenancy"`
}
