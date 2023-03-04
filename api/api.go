package api

import (
	filesV1 "github.com/turao/topics/api/files/v1"
	moviesV1 "github.com/turao/topics/api/movies/v1"
	usersV1 "github.com/turao/topics/api/users/v1"
)

type API interface {
	moviesV1.Movies
	filesV1.Files
	usersV1.Users
}
