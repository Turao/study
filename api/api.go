package api

import (
	moviesV1 "github.com/turao/topics/api/movies/v1"
	usersV1 "github.com/turao/topics/api/users/v1"
)

type API interface {
	moviesV1.Movies
	moviesV1.Files
	usersV1.Users
}
