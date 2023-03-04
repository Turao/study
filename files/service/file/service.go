package file

import (
	"context"
	"errors"
	"log"

	apiV1 "github.com/turao/topics/api/files/v1"
	eventsV1 "github.com/turao/topics/events/files/v1"
	"github.com/turao/topics/files/entity/file"
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

func (svc *service) DownloadFile(ctx context.Context, req apiV1.DownloadFileRequest) (apiV1.DownloadFileResponse, error) {
	log.Println("Downloading file from:", req.URI)

	// todo: download logic
	// download to disk
	// create file
	filecfg, errs := file.NewConfig(
		file.WithMovieID(movie.ID(req.MovieID)),
		file.WithURI("movie-stored-here"),
		file.WithSize(10000),
	)
	if len(errs) > 0 {
		return apiV1.DownloadFileResponse{}, errors.Join(errs...)
	}
	file := file.NewFile(filecfg)
	err := svc.fileRepository.Save(ctx, file)
	if err != nil {
		return apiV1.DownloadFileResponse{}, err
	}

	log.Println(eventsV1.NewFileDownloaded(file))

	return apiV1.DownloadFileResponse{}, nil
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
