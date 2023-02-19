package movie

type Model struct {
	ID    string `json:"id"`
	Title string `json:"title"`

	Tenancy string `json:"tenancy"`

	CreatedAt int64  `json:"created_at"`
	DeletedAt *int64 `json:"deleted_at"`
}
