package notification

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/turao/topics/notifications/entity/notification"
)

type repository struct {
	database *sqlx.DB
}

// NewRepository creates a new notification repository
func NewRepository(database *sqlx.DB) (*repository, error) {
	if database == nil {
		return nil, errors.New("database connection is nil")
	}

	return &repository{
		database: database,
	}, nil
}

// Save saves a notification
func (r *repository) Save(ctx context.Context, notification notification.Notification) error {
	model, err := ToModel(notification)
	if err != nil {
		return err
	}

	_, err = r.database.NamedExecContext(
		ctx,
		`INSERT INTO notifications (id, type, recipient, subject, content, created_at, external_reference_id) 
		VALUES (:id, :type, :recipient, :subject, :content, :created_at, :external_reference_id)
		ON CONFLICT (id) DO UPDATE SET
			type = :type,
			recipient = :recipient,
			subject = :subject,
			content = :content,
			created_at = :created_at,
			external_reference_id = :external_reference_id`,
		model,
	)
	return err
}

// FindByID finds a notification by its ID
func (r *repository) FindByID(ctx context.Context, id string) (*notification.Notification, error) {
	var model Model
	err := r.database.GetContext(
		ctx,
		&model,
		"SELECT * FROM notifications WHERE id = $1 ORDER BY created_at DESC LIMIT 1",
		id,
	)
	if err != nil {
		return nil, nil
	}

	return ToEntity(model)
}
