package v1

import (
	"context"
)

type Movies interface {
	RegisterMovie(ctx context.Context, req RegisterMovieRequest) (RegisterMovieResponse, error)
	ListMovies(ctx context.Context, req ListMoviesRequest) (ListMoviesResponse, error)
	GetMovie(ctx context.Context, req GetMovieRequest) (GetMovieResponse, error)
}

type Movie struct {
	ID        string `json:"id"`
	Tenancy   string `json:"tenancy"`
	CreatedAt string `json:"createdAt"`
	DeletedAt string `json:"deletedAt"`
}

type RegisterMovieRequest struct {
	Tenancy string `json:"tenancy"`
}

type RegisterMovieResponse struct{}

type ListMoviesRequest struct{}

type ListMoviesResponse struct {
	Movies []Movie `json:"movies"`
}

type GetMovieRequest struct {
	MovieID string `json:"movieId"`
}

type GetMovieResponse struct {
	Movie Movie `json:"movie"`
}
