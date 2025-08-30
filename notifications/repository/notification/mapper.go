package notification

import (
	"encoding/json"

	"github.com/turao/topics/notifications/entity/notification"
)

// ToModel converts a Notification entity to a Model
func ToModel(notification notification.Notification) (*Model, error) {
	content, err := json.Marshal(notification.Content)
	if err != nil {
		return nil, err
	}

	return &Model{
		ID:                  notification.ID,
		Type:                notification.Type,
		Recipient:           notification.Recipient,
		Subject:             notification.Subject,
		Content:             string(content),
		CreatedAt:           notification.CreatedAt,
		ExternalReferenceID: notification.ExternalReferenceID,
	}, nil
}

// ToEntity converts a Model to a Notification entity
func ToEntity(model Model) (*notification.Notification, error) {
	var content map[string]interface{}
	err := json.Unmarshal([]byte(model.Content), &content)
	if err != nil {
		return nil, err
	}

	return &notification.Notification{
		ID:                  model.ID,
		Type:                model.Type,
		Recipient:           model.Recipient,
		Subject:             model.Subject,
		Content:             content,
		CreatedAt:           model.CreatedAt,
		ExternalReferenceID: model.ExternalReferenceID,
	}, nil
}
