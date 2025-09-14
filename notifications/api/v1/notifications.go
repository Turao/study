package api

import "context"

type Notifications interface {
	SendNotification(ctx context.Context, req SendNotificationRequest) (SendNotificationResponse, error)
}

type SendNotificationRequest struct {
	Recipient        string                 `json:"recipient"`
	NotificationType string                 `json:"notification_type"`
	Metadata         map[string]interface{} `json:"metadata"`

	// notification specific fields
	PaymentSettled *PaymentSettled `json:"payment_settled"`
}

type PaymentSettled struct {
	PaymentID string `json:"payment_id"`
}

type SendNotificationResponse struct {
	ID string `json:"id"`
}
