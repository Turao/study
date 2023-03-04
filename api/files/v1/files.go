package v1

import (
	"context"
)

type Files interface {
	ListFilesByMovie(ctx context.Context, req ListFilesByMovieRequest) (ListFilesByMovieResponse, error)

	DownloadFile(ctx context.Context, req DownloadFileRequest) (DownloadFileResponse, error)
}

type File struct {
	ID      string `json:"id"`
	MovieID string `json:"movieId"`
	URI     string `json:"uri"`
	Size    int64  `json:"size"`

	Tenancy   string `json:"tenancy"`
	CreatedAt string `json:"createdAt"`
	DeletedAt string `json:"deletedAt"`
}

type ListFilesByMovieRequest struct {
	MovieID string `json:"movieId"`
}

type ListFilesByMovieResponse struct {
	Files []File `json:"files"`
}

type DownloadFileRequest struct {
	URI     string `json:"uri"`
	MovieID string `json:"movieId"`
}

type DownloadFileResponse struct {
	URI string `json:"uri"`
}
