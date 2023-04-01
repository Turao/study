package v1

import "context"

type Messages interface {
	SendMessage(ctx context.Context, req SendMessageRequest) (SendMessageResponse, error)
}

type SendMessageRequest struct {
	Channel string
	Content string
}

type SendMessageResponse struct{}
