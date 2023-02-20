package file

type Model struct {
	ID   string `json:"id"`
	URI  string `json:"uri"`
	Size int64  `json:"size"`

	Tenancy string `json:"tenancy"`

	CreatedAt int64  `json:"created_at"`
	DeletedAt *int64 `json:"deleted_at"`
}
