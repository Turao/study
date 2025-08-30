package confirmation

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	apiV1 "github.com/turao/topics/notifications/api/v1"
	"github.com/turao/topics/notifications/entity/notification"
)

type builder struct{}

// Newbuilder creates a new notification builder
func NewBuilder() *builder {
	return &builder{}
}

func (b *builder) NotificationType() string {
	return "confirmation"
}

// SendNotification sends a notification
func (b *builder) BuildNotification(ctx context.Context, request apiV1.SendNotificationRequest) (*notification.Notification, error) {
	return &notification.Notification{
		ID:        uuid.Must(uuid.NewV4()).String(),
		Type:      b.NotificationType(),
		Recipient: "john@doe.com",
		Subject:   "Payment confirmed!",
		Content: map[string]interface{}{
			"price_e5":                      99900,
			"discount_percentage":           20,
			"payment_method_display_name":   "Apple Pay",
			"payment_method_display_number": "**** **** *234",
			"billing_date_yyyy_mm_dd":       "2025-08-30",
		},
		CreatedAt:           time.Now(),
		ExternalReferenceID: uuid.Must(uuid.NewV4()).String(),
	}, nil
}
