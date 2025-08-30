package builder

import (
	"context"
	"errors"
	"log"

	apiV1 "github.com/turao/topics/notifications/api/v1"
	"github.com/turao/topics/notifications/entity/notification"
)

type NotificationBuilder interface {
	NotificationType() string
	BuildNotification(ctx context.Context, request apiV1.SendNotificationRequest) (*notification.Notification, error)
}

type builder struct {
	builders map[string]NotificationBuilder
}

// NewBuilder creates a new notification builder
func NewBuilder(notificationBuilders ...NotificationBuilder) *builder {
	builders := make(map[string]NotificationBuilder, len(notificationBuilders))
	for _, notificationBuilder := range notificationBuilders {
		builders[notificationBuilder.NotificationType()] = notificationBuilder
	}

	return &builder{
		builders: builders,
	}
}

// SendNotification sends a notification
func (s *builder) BuildNotification(ctx context.Context, request apiV1.SendNotificationRequest) (*notification.Notification, error) {
	builder, ok := s.builders[request.NotificationType]
	if !ok {
		return nil, errors.New("builder not found for notification type")
	}

	log.Println("building notification", request.NotificationType)
	return builder.BuildNotification(ctx, request)
}
