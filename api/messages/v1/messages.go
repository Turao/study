package v1

import "context"

type Messages interface {
	SendMessage(ctx context.Context, req SendMessageRequest) (SendMessageResponse, error)
}

type SendMessageRequest struct {
	AuthorID  string `json:"authorId"`
	ChannelID string `json:"channelId"`
	Content   string `json:"content"`
}

type SendMessageResponse struct{}
