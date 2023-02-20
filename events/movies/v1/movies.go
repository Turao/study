package v1

type MovieRegistered struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	URI        string `json:"uri"`
	Downloaded bool   `json:"downloaded"`
	Tenancy    string `json:"tenancy"`
	CreatedAt  string `json:"createdAt"`
	DeletedAt  string `json:"deletedAt"`
}
