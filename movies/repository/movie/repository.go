package movie

import (
	"context"
	"errors"

	"github.com/turao/topics/movies/entity/movie"
)

var (
	ErrNotFound = errors.New("not found")
)

type repository struct {
	movies map[string]*Model
}

func NewRepository() (*repository, error) {
	return &repository{
		movies: make(map[string]*Model),
	}, nil
}

func (r *repository) FindAll(ctx context.Context) ([]movie.Movie, error) {
	var movies []movie.Movie
	for _, model := range r.movies {
		movie, err := ToEntity(*model)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (r *repository) FindByID(ctx context.Context, movieID movie.ID) (movie.Movie, error) {
	model, found := r.movies[movieID.String()]
	if !found {
		return nil, ErrNotFound
	}
	return ToEntity(*model)
}

func (r *repository) Save(ctx context.Context, movie movie.Movie) error {
	model, err := ToModel(movie)
	if err != nil {
		return err
	}

	r.movies[movie.ID().String()] = model
	return nil
}
