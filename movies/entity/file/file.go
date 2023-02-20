package file

import (
	"time"

	"github.com/turao/topics/metadata"
	"github.com/turao/topics/movies/entity/movie"
)

type ID string

func (id ID) String() string {
	return string(id)
}

type File interface {
	ID() ID
	Movie() movie.ID
	URI() string
	Size() int64

	Delete()
	metadata.Auditable
	metadata.MultiTenant
}

type file struct {
	cfg config
}

var _ File = (*file)(nil)

func NewFile(cfg config) *file {
	return &file{
		cfg: cfg,
	}
}

func (f *file) ID() ID {
	return f.cfg.id
}

func (f *file) Movie() movie.ID {
	return f.cfg.movieID
}

func (f *file) URI() string {
	return f.cfg.uri
}

func (f *file) Size() int64 {
	return f.cfg.size
}

func (f *file) Tenancy() metadata.Tenancy {
	return f.cfg.tenancy
}

func (f *file) CreatedAt() time.Time {
	return f.cfg.createdAt
}

func (f *file) DeletedAt() *time.Time {
	return f.cfg.deletedAt
}

func (f *file) Delete() {
	if f.DeletedAt() == nil {
		now := time.Now()
		f.cfg.deletedAt = &now
	}
}
