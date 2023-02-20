package file

import (
	"context"

	apiV1 "github.com/turao/topics/api/movies/v1"
	"github.com/turao/topics/movies/entity/file"
	"github.com/turao/topics/movies/entity/movie"
)

type FileRepository interface {
	FindByMovieID(ctx context.Context, movieID movie.ID) ([]file.File, error)
	FindByID(ctx context.Context, fileID file.ID) (file.File, error)
	Save(ctx context.Context, file file.File) error
}

type service struct {
	fileRepository FileRepository
}

var _ apiV1.Files = (*service)(nil)

func NewService(
	fileRepository FileRepository,
) (*service, error) {
	return &service{
		fileRepository: fileRepository,
	}, nil
}

func (svc *service) ListFilesByMovie(ctx context.Context, req apiV1.ListFilesByMovieRequest) (apiV1.ListFilesByMovieResponse, error) {
	files, err := svc.fileRepository.FindByMovieID(ctx, movie.ID(req.MovieID))
	if err != nil {
		return apiV1.ListFilesByMovieResponse{}, err
	}

	res := apiV1.ListFilesByMovieResponse{Files: make([]apiV1.File, 0)}
	for _, file := range files {
		res.Files = append(res.Files, apiV1.File{
			ID:        file.ID().String(),
			MovieID:   file.Movie().String(),
			URI:       file.URI(),
			Size:      file.Size(),
			Tenancy:   file.Tenancy().String(),
			CreatedAt: file.CreatedAt().String(),
			// DeletedAt: file.DeletedAt().String(),
		})
	}
	return res, nil
}
