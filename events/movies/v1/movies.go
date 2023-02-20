package v1

import "github.com/turao/topics/movies/entity/movie"

type MovieRegistered struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	URI        string `json:"uri"`
	Downloaded bool   `json:"downloaded"`
	Tenancy    string `json:"tenancy"`
	CreatedAt  string `json:"createdAt"`
	DeletedAt  string `json:"deletedAt"`
}

func NewMovieRegistered(movie movie.Movie) *MovieRegistered {
	event := &MovieRegistered{
		ID:         movie.ID().String(),
		Title:      movie.Title(),
		URI:        movie.URI(),
		Downloaded: movie.Downloaded(),
		Tenancy:    movie.Tenancy().String(),
		CreatedAt:  movie.CreatedAt().String(),
	}

	if movie.DeletedAt() != nil {
		event.DeletedAt = movie.DeletedAt().String()
	}

	return event
}
