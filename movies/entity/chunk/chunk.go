package chunk

import (
	"time"

	"github.com/turao/topics/metadata"
	"github.com/turao/topics/movies/entity/movie"
)

type ID string

func (id ID) String() string {
	return string(id)
}

type Chunk interface {
	ID() ID
	Movie() movie.ID
	URI() string
	Encoded() bool
	Content() []byte

	metadata.Auditable
}

type chunk struct {
	cfg config
}

var _ Chunk = (*chunk)(nil)

func NewChunk(cfg config) *chunk {
	return &chunk{
		cfg: cfg,
	}
}

func (c *chunk) ID() ID {
	return c.cfg.id
}

func (c *chunk) Movie() movie.ID {
	return c.cfg.movieID
}

func (c *chunk) URI() string {
	return c.cfg.uri
}

func (c *chunk) Content() []byte {
	return c.cfg.content
}

func (c *chunk) Encoded() bool {
	return c.cfg.encoded
}

func (c *chunk) CreatedAt() time.Time {
	return c.cfg.createdAt
}

func (c *chunk) DeletedAt() *time.Time {
	return c.cfg.deletedAt
}
