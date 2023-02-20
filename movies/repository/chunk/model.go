package chunk

type Model struct {
	ID      string `json:"id"`
	MovieID string `json:""movie_id`
	URI     string `json:"uri"`
	Encoded bool   `json:"encoded"`
	Content []byte `json:"content"`

	CreatedAt int64  `json:"created_at"`
	DeletedAt *int64 `json:"deleted_at"`
}
