package v1

import (
	"context"
)

type Movies interface {
	RegisterMovie(ctx context.Context, req RegisterMovieRequest) (RegisterMovieResponse, error)
	ListMovies(ctx context.Context, req ListMoviesRequest) (ListMoviesResponse, error)
	GetMovie(ctx context.Context, req GetMovieRequest) (GetMovieResponse, error)
	DeleteMovie(ctx context.Context, req DeleteMovieRequest) (DeleteMovieResponse, error)
	MarkAsDownloaded(ctx context.Context, req MarkAsDownloadedRequest) (MarkAsDownloadedResponse, error)
}

type Movie struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	URI        string `json:"uri"`
	Downloaded bool   `json:"downloaded"`
	Tenancy    string `json:"tenancy"`
	CreatedAt  string `json:"createdAt"`
	DeletedAt  string `json:"deletedAt"`
}

type RegisterMovieRequest struct {
	Title   string `json:"title"`
	URI     string `json:"uri"`
	Tenancy string `json:"tenancy"`
}

type RegisterMovieResponse struct {
	ID string `json:"id"`
}

type ListMoviesRequest struct{}

type ListMoviesResponse struct {
	Movies []Movie `json:"movies"`
}

type GetMovieRequest struct {
	ID string `json:"id"`
}

type GetMovieResponse struct {
	Movie Movie `json:"movie"`
}

type DeleteMovieRequest struct {
	ID string `json:"id"`
}

type DeleteMovieResponse struct{}

type MarkAsDownloadedRequest struct {
	ID string `json:"id"`
}

type MarkAsDownloadedResponse struct{}
