package chunk

import (
	"errors"
	"time"

	"github.com/turao/topics/movies/entity/chunk"
	"github.com/turao/topics/movies/entity/movie"
)

func ToModel(chunk chunk.Chunk) (*Model, error) {
	model := &Model{
		ID:        chunk.ID().String(),
		MovieID:   chunk.Movie().String(),
		URI:       chunk.URI(),
		Encoded:   chunk.Encoded(),
		Content:   chunk.Content(),
		CreatedAt: chunk.CreatedAt().UnixMilli(),
	}

	if chunk.DeletedAt() != nil {
		ts := chunk.DeletedAt().UnixMilli()
		model.DeletedAt = &ts
	}

	return model, nil
}

func ToEntity(model Model) (chunk.Chunk, error) {
	var deletedAt *time.Time
	if model.DeletedAt != nil {
		ts := time.UnixMilli(*model.DeletedAt)
		deletedAt = &ts
	}

	chunkcfg, errs := chunk.NewConfig(
		chunk.WithID(chunk.ID(model.ID)),
		chunk.WithMovieID(movie.ID(model.MovieID)),
		chunk.WithURI(model.URI),
		chunk.WithEncoded(model.Encoded),
		chunk.WithCreatedAt(time.UnixMilli(model.CreatedAt)),
		chunk.WithDeletedAt(deletedAt),
	)
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	entity := chunk.NewChunk(chunkcfg)
	return entity, nil
}
