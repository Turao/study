package api

import (
	channelsV1 "github.com/turao/topics/api/channels/v1"
	messagesV1 "github.com/turao/topics/api/messages/v1"
	usersV1 "github.com/turao/topics/api/users/v1"
)

type API interface {
	usersV1.Users
	channelsV1.Channels
	messagesV1.Messages
}
