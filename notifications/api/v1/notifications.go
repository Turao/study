package api

import "context"

type Notifications interface {
	SendNotification(ctx context.Context, req SendNotificationRequest) (SendNotificationResponse, error)
}

type SendNotificationRequest struct {
	NotificationType string `json:"type"`
}

type SendNotificationResponse struct {
	ID string `json:"id"`
}
