package notification

import (
	"context"

	apiV1 "github.com/turao/topics/notifications/api/v1"
	"github.com/turao/topics/notifications/entity/notification"
)

type NotificationRepository interface {
	Save(ctx context.Context, notification notification.Notification) error
	FindByID(ctx context.Context, id string) (*notification.Notification, error)
}

type NotificationBuilder interface {
	BuildNotification(ctx context.Context, request apiV1.SendNotificationRequest) (*notification.Notification, error)
}

type NotificationSender interface {
	SendNotification(ctx context.Context, notification notification.Notification) error
}

type service struct {
	notificationBuilder    NotificationBuilder
	notificationRepository NotificationRepository
	notificationSender     NotificationSender
}

var _ apiV1.Notifications = (*service)(nil)

// NewService creates a new notification service
func NewService(
	notificationBuilder NotificationBuilder,
	notificationRepository NotificationRepository,
	notificationSender NotificationSender,
) *service {
	return &service{
		notificationBuilder:    notificationBuilder,
		notificationRepository: notificationRepository,
		notificationSender:     notificationSender,
	}
}

// SendNotification sends a notification
func (s *service) SendNotification(ctx context.Context, request apiV1.SendNotificationRequest) (apiV1.SendNotificationResponse, error) {
	notification, err := s.notificationBuilder.BuildNotification(ctx, request)
	if err != nil {
		return apiV1.SendNotificationResponse{}, err
	}

	if notification == nil {
		return apiV1.SendNotificationResponse{}, nil
	}

	err = s.notificationRepository.Save(ctx, *notification)
	if err != nil {
		return apiV1.SendNotificationResponse{}, err
	}

	err = s.notificationSender.SendNotification(ctx, *notification)
	if err != nil {
		return apiV1.SendNotificationResponse{}, err
	}

	return apiV1.SendNotificationResponse{}, nil
}
