package sender

import (
	"context"
	"log"

	"github.com/turao/topics/notifications/entity/notification"
)

type sender struct{}

// NewSender creates a new notification sender
func NewSender() *sender {
	return &sender{}
}

// SendNotification sends a notification
func (s *sender) SendNotification(ctx context.Context, notification notification.Notification) error {
	log.Println("sending notification", notification)
	return nil
}
