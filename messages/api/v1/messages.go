package v1

import (
	"context"
	"time"
)

type Messages interface {
	SendMessage(ctx context.Context, req SendMessageRequest) (SendMessageResponse, error)
	GetMessages(ctx context.Context, req GetMessagesRequest) (GetMessagesResponse, error)
	DeleteMessage(ctx context.Context, req DeleteMessageRequest) (DeleteMessageResponse, error)
	GetMessageStream(ctx context.Context, req GetMessageStreamRequest) (GetMessageStreamResponse, error)
}

type SendMessageRequest struct {
	AuthorID  string `json:"authorId"`
	ChannelID string `json:"channelId"`
	Content   string `json:"content"`
	Tenancy   string `json:"tenancy"`
}

type SendMessageResponse struct{}

type GetMessagesRequest struct {
	ChannelID string `json:"channelId"`
}

type GetMessagesResponse struct {
	Messages []MessageInfo `json:"messages"`
}

type DeleteMessageRequest struct {
	ChannelID string `json:"channelId"`
	MessageID string `json:"messageId"`
}

type DeleteMessageResponse struct{}

type GetMessageStreamRequest struct {
	ChannelID string `json:"channelId"`
}

type GetMessageStreamResponse struct {
	Messages <-chan MessageInfo `json:"messages"`
}

type MessageInfo struct {
	ID        string     `json:"id"`
	Version   uint32     `json:"version"`
	Author    string     `json:"author"`
	Content   string     `json:"content"`
	Tenancy   string     `json:"tenancy"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
