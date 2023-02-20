package movie

import (
	"context"
	"errors"

	v1 "github.com/turao/topics/api/movies/v1"
	"github.com/turao/topics/metadata"
	"github.com/turao/topics/movies/entity/file"
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
	fileRepository  FileRepository
}

var _ v1.Movies = (*service)(nil)

func NewService(
	movieRepository MovieRepository,
	fileRepository FileRepository,
) (*service, error) {
	return &service{
		movieRepository: movieRepository,
		fileRepository:  fileRepository,
	}, nil
}

func (svc *service) DeleteMovie(ctx context.Context, req v1.DeleteMovieRequest) (v1.DeleteMovieResponse, error) {
	movie, err := svc.movieRepository.FindByID(ctx, movie.ID(req.ID))
	if err != nil {
		return v1.DeleteMovieResponse{}, err
	}

	movie.Delete()
	err = svc.movieRepository.Save(ctx, movie)
	if err != nil {
		return v1.DeleteMovieResponse{}, err
	}

	return v1.DeleteMovieResponse{}, nil
}

func (svc *service) GetMovie(ctx context.Context, req v1.GetMovieRequest) (v1.GetMovieResponse, error) {
	movie, err := svc.movieRepository.FindByID(ctx, movie.ID(req.ID))
	if err != nil {
		return v1.GetMovieResponse{}, err
	}

	return v1.GetMovieResponse{
		Movie: v1.Movie{
			ID:        movie.ID().String(),
			Tenancy:   movie.Tenancy().String(),
			CreatedAt: movie.CreatedAt().String(),
			// DeletedAt: movie.DeletedAt().String(),
		},
	}, nil
}

func (svc *service) ListMovies(ctx context.Context, req v1.ListMoviesRequest) (v1.ListMoviesResponse, error) {
	movies, err := svc.movieRepository.FindAll(ctx)
	if err != nil {
		return v1.ListMoviesResponse{}, err
	}

	res := v1.ListMoviesResponse{Movies: make([]v1.Movie, 0)}
	for _, movie := range movies {
		res.Movies = append(res.Movies, v1.Movie{
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

func (svc *service) RegisterMovie(ctx context.Context, req v1.RegisterMovieRequest) (v1.RegisterMovieResponse, error) {
	moviecfg, errs := movie.NewConfig(
		movie.WithTitle(req.Title),
		movie.WithURI(req.URI),
		movie.WithTenancy(metadata.Tenancy(req.Tenancy)),
	)
	if len(errs) > 0 {
		return v1.RegisterMovieResponse{}, errors.Join(errs...)
	}

	movie := movie.NewMovie(moviecfg)
	err := svc.movieRepository.Save(ctx, movie)
	if err != nil {
		return v1.RegisterMovieResponse{}, err
	}

	return v1.RegisterMovieResponse{
		ID: movie.ID().String(),
	}, nil
}

func (svc *service) DownloadMovie(ctx context.Context, req v1.DownloadMovieRequest) (v1.DownloadMovieResponse, error) {
	movie, err := svc.movieRepository.FindByID(ctx, movie.ID(req.ID))
	if err != nil {
		return v1.DownloadMovieResponse{}, err
	}

	if movie.Downloaded() {
		return v1.DownloadMovieResponse{}, nil // skip
	}

	// todo: download logic
	// download to disk
	// create file
	filecfg, errs := file.NewConfig(
		file.WithMovieID(movie.ID()),
		file.WithURI("movie-stored-here"),
		file.WithSize(10000),
	)
	if len(errs) > 0 {
		return v1.DownloadMovieResponse{}, errors.Join(errs...)
	}
	file := file.NewFile(filecfg)
	err = svc.fileRepository.Save(ctx, file)
	if err != nil {
		return v1.DownloadMovieResponse{}, err
	}

	movie.MarkAsDownloaded()
	err = svc.movieRepository.Save(ctx, movie)
	if err != nil {
		return v1.DownloadMovieResponse{}, err
	}

	return v1.DownloadMovieResponse{}, nil
}
