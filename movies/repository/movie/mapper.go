package movie

import (
	"errors"
	"time"

	"github.com/turao/topics/metadata"
	"github.com/turao/topics/movies/entity/movie"
)

func ToModel(movie movie.Movie) (*Model, error) {
	model := &Model{
		ID:         movie.ID().String(),
		Title:      movie.Title(),
		URI:        movie.URI(),
		Downloaded: movie.Downloaded(),
		Tenancy:    movie.Tenancy().String(),
		CreatedAt:  movie.CreatedAt().UnixMilli(),
	}

	if movie.DeletedAt() != nil {
		ts := movie.DeletedAt().UnixMilli()
		model.DeletedAt = &ts
	}

	return model, nil
}

func ToEntity(model Model) (movie.Movie, error) {
	var deletedAt *time.Time
	if model.DeletedAt != nil {
		ts := time.UnixMilli(*model.DeletedAt)
		deletedAt = &ts
	}

	moviecfg, errs := movie.NewConfig(
		movie.WithID(movie.ID(model.ID)),
		movie.WithTitle(model.Title),
		movie.WithURI(model.URI),
		movie.WithDownloaded(model.Downloaded),
		movie.WithTenancy(metadata.Tenancy(model.Tenancy)),
		movie.WithCreatedAt(time.UnixMilli(model.CreatedAt)),
		movie.WithDeletedAt(deletedAt),
	)
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	entity := movie.NewMovie(moviecfg)
	return entity, nil
}
