package file

import (
	"context"
	"errors"

	"github.com/turao/topics/files/entity/file"
	"github.com/turao/topics/movies/entity/movie"
)

var (
	ErrNotFound = errors.New("not found")
)

type repository struct {
	filesByID      map[string]*Model
	filesByMovieID map[string][]*Model
}

func NewRepository() (*repository, error) {
	return &repository{
		filesByID:      make(map[string]*Model),
		filesByMovieID: make(map[string][]*Model),
	}, nil
}

func (r *repository) FindByMovieID(ctx context.Context, movieID movie.ID) ([]file.File, error) {
	filesByMovieID, found := r.filesByMovieID[movieID.String()]
	if !found {
		return nil, ErrNotFound
	}

	var files []file.File
	for _, model := range filesByMovieID {
		file, err := ToEntity(*model)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}

func (r *repository) FindByID(ctx context.Context, fileID file.ID) (file.File, error) {
	model, found := r.filesByID[fileID.String()]
	if !found {
		return nil, ErrNotFound
	}
	return ToEntity(*model)
}

func (r *repository) Save(ctx context.Context, file file.File) error {
	model, err := ToModel(file)
	if err != nil {
		return err
	}

	r.filesByID[file.ID().String()] = model

	_, found := r.filesByMovieID[file.Movie().String()]
	if !found {
		r.filesByMovieID[file.Movie().String()] = make([]*Model, 0)
	}
	r.filesByMovieID[file.Movie().String()] = append(
		r.filesByMovieID[file.Movie().String()],
		model,
	)
	return nil
}
