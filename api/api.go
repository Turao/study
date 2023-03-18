package api

import (
	usersV1 "github.com/turao/topics/api/users/v1"
)

type API interface {
	usersV1.Users
}
