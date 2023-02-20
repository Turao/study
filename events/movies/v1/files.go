package v1

type FileCreated struct {
	ID      string `json:"id"`
	MovieID string `json:"movieId"`
	URI     string `json:"uri"`
	Size    int64  `json:"size"`

	Tenancy   string `json:"tenancy"`
	CreatedAt string `json:"createdAt"`
	DeletedAt string `json:"deletedAt"`
}
