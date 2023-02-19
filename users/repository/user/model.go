package user

type Model struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`

	Tenancy string `json:"tenancy"`

	CreatedAt int64  `json:"created_at"`
	DeletedAt *int64 `json:"deleted_at"`
}
