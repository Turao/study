package v1

import "context"

type Messages interface {
	SendMessage(ctx context.Context, req SendMessageRequest) (SendMessageResponse, error)
}

type SendMessageRequest struct {
	Author  string
	Channel string
	Content string
}

type SendMessageResponse struct{}
