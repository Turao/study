package chunk

import (
	"context"
	"errors"

	"github.com/turao/topics/movies/entity/chunk"
	"github.com/turao/topics/movies/entity/movie"
)

var (
	ErrNotFound = errors.New("not found")
)

type repository struct {
	chunksByID      map[string]*Model
	chunksByMovieID map[string][]*Model
}

func NewRepository() (*repository, error) {
	return &repository{
		chunksByID:      make(map[string]*Model),
		chunksByMovieID: make(map[string][]*Model),
	}, nil
}

func (r *repository) FindByMovieID(ctx context.Context, movieID movie.ID) ([]chunk.Chunk, error) {
	movieChunks, found := r.chunksByMovieID[movieID.String()]
	if !found {
		return nil, ErrNotFound
	}

	var chunks []chunk.Chunk
	for _, model := range movieChunks {
		chunk, err := ToEntity(*model)
		if err != nil {
			return nil, err
		}
		chunks = append(chunks, chunk)
	}

	return chunks, nil
}

func (r *repository) FindByID(ctx context.Context, chunkID chunk.ID) (chunk.Chunk, error) {
	model, found := r.chunksByID[chunkID.String()]
	if !found {
		return nil, ErrNotFound
	}
	return ToEntity(*model)
}

func (r *repository) Save(ctx context.Context, chunk chunk.Chunk) error {
	model, err := ToModel(chunk)
	if err != nil {
		return err
	}

	r.chunksByID[chunk.ID().String()] = model

	_, found := r.chunksByMovieID[chunk.Movie().String()]
	if !found {
		r.chunksByMovieID[chunk.Movie().String()] = make([]*Model, 0)
	}
	r.chunksByMovieID[chunk.Movie().String()] = append(
		r.chunksByMovieID[chunk.Movie().String()],
		model,
	)

	return nil
}
