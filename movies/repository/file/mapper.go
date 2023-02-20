package file

import (
	"errors"
	"time"

	"github.com/turao/topics/metadata"
	"github.com/turao/topics/movies/entity/file"
	"github.com/turao/topics/movies/entity/movie"
)

func ToModel(file file.File) (*Model, error) {
	model := &Model{
		ID:        file.ID().String(),
		MovieID:   file.Movie().String(),
		URI:       file.URI(),
		Size:      file.Size(),
		Tenancy:   file.Tenancy().String(),
		CreatedAt: file.CreatedAt().UnixMilli(),
	}

	if file.DeletedAt() != nil {
		ts := file.DeletedAt().UnixMilli()
		model.DeletedAt = &ts
	}

	return model, nil
}

func ToEntity(model Model) (file.File, error) {
	var deletedAt *time.Time
	if model.DeletedAt != nil {
		ts := time.UnixMilli(*model.DeletedAt)
		deletedAt = &ts
	}

	filecfg, errs := file.NewConfig(
		file.WithID(file.ID(model.ID)),
		file.WithMovieID(movie.ID(model.MovieID)),
		file.WithURI(model.URI),
		file.WithSize(model.Size),
		file.WithTenancy(metadata.Tenancy(model.Tenancy)),
		file.WithCreatedAt(time.UnixMilli(model.CreatedAt)),
		file.WithDeletedAt(deletedAt),
	)
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	entity := file.NewFile(filecfg)
	return entity, nil
}
