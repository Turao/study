package movie

import (
	"context"
	"errors"
	"log"

	apiV1 "github.com/turao/topics/api/movies/v1"
	eventsV1 "github.com/turao/topics/events/movies/v1"
	"github.com/turao/topics/files/entity/file"
	"github.com/turao/topics/metadata"
	"github.com/turao/topics/movies/entity/movie"
)

type MovieRepository interface {
	Save(ctx context.Context, movie movie.Movie) error
	FindByID(ctx context.Context, movieID movie.ID) (movie.Movie, error)
	FindAll(ctx context.Context) ([]movie.Movie, error)
}

type FileRepository interface {
	FindByMovieID(ctx context.Context, movieID movie.ID) ([]file.File, error)
	FindByID(ctx context.Context, fileID file.ID) (file.File, error)
	Save(ctx context.Context, file file.File) error
}

type service struct {
	movieRepository MovieRepository
}

var _ apiV1.Movies = (*service)(nil)

func NewService(
	movieRepository MovieRepository,
) (*service, error) {
	return &service{
		movieRepository: movieRepository,
	}, nil
}

func (svc *service) DeleteMovie(ctx context.Context, req apiV1.DeleteMovieRequest) (apiV1.DeleteMovieResponse, error) {
	movie, err := svc.movieRepository.FindByID(ctx, movie.ID(req.ID))
	if err != nil {
		return apiV1.DeleteMovieResponse{}, err
	}

	movie.Delete()
	err = svc.movieRepository.Save(ctx, movie)
	if err != nil {
		return apiV1.DeleteMovieResponse{}, err
	}

	return apiV1.DeleteMovieResponse{}, nil
}

func (svc *service) GetMovie(ctx context.Context, req apiV1.GetMovieRequest) (apiV1.GetMovieResponse, error) {
	movie, err := svc.movieRepository.FindByID(ctx, movie.ID(req.ID))
	if err != nil {
		return apiV1.GetMovieResponse{}, err
	}

	return apiV1.GetMovieResponse{
		Movie: apiV1.Movie{
			ID:        movie.ID().String(),
			Tenancy:   movie.Tenancy().String(),
			CreatedAt: movie.CreatedAt().String(),
			// DeletedAt: movie.DeletedAt().String(),
		},
	}, nil
}

func (svc *service) ListMovies(ctx context.Context, req apiV1.ListMoviesRequest) (apiV1.ListMoviesResponse, error) {
	movies, err := svc.movieRepository.FindAll(ctx)
	if err != nil {
		return apiV1.ListMoviesResponse{}, err
	}

	res := apiV1.ListMoviesResponse{Movies: make([]apiV1.Movie, 0)}
	for _, movie := range movies {
		res.Movies = append(res.Movies, apiV1.Movie{
			ID:         movie.ID().String(),
			Title:      movie.Title(),
			URI:        movie.URI(),
			Downloaded: movie.Downloaded(),
			Tenancy:    movie.Tenancy().String(),
			CreatedAt:  movie.CreatedAt().String(),
			// DeletedAt: movie.DeletedAt().String(),
		})
	}
	return res, nil
}

func (svc *service) RegisterMovie(ctx context.Context, req apiV1.RegisterMovieRequest) (apiV1.RegisterMovieResponse, error) {
	moviecfg, errs := movie.NewConfig(
		movie.WithTitle(req.Title),
		movie.WithURI(req.URI),
		movie.WithTenancy(metadata.Tenancy(req.Tenancy)),
	)
	if len(errs) > 0 {
		return apiV1.RegisterMovieResponse{}, errors.Join(errs...)
	}

	movie := movie.NewMovie(moviecfg)
	err := svc.movieRepository.Save(ctx, movie)
	if err != nil {
		return apiV1.RegisterMovieResponse{}, err
	}

	log.Println(eventsV1.NewMovieRegistered(movie))

	return apiV1.RegisterMovieResponse{
		ID: movie.ID().String(),
	}, nil
}
