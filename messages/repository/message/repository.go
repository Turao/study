package message

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"github.com/turao/topics/channels/entity/channel"
	"github.com/turao/topics/messages/entity/message"
)

type repository struct {
	database gocqlx.Session
}

func NewRepository(database gocqlx.Session) (*repository, error) {
	return &repository{
		database: database,
	}, nil
}

func (r *repository) Save(ctx context.Context, message message.Message) error {
	msg, err := ToModel(message)
	if err != nil {
		return err
	}

	err = r.database.Query(_table.Insert()).WithContext(ctx).BindStruct(msg).ExecRelease()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindOne(
	ctx context.Context,
	channelID channel.ID,
	messageID message.ID,
) (message.Message, error) {
	model := Model{}

	err := r.database.
		Query(_table.Select()).
		BindMap(map[string]interface{}{
			"id":      messageID.String(),
			"channel": channelID.String(),
		}).
		WithContext(ctx).
		Get(&model)
	if err != nil {
		return nil, err
	}

	return ToEntity(model)
}

func (r *repository) ListAllByChannelID(ctx context.Context, channelID channel.ID) ([]message.Message, error) {
	models := []Model{}
	err := r.database.Query(_table.SelectAll()).WithContext(ctx).SelectRelease(&models)
	if err != nil {
		return nil, err
	}

	messages := []message.Message{}
	for _, model := range models {
		msg, err := ToEntity(model)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

func (r *repository) StreamAllByChannelID(ctx context.Context, channelID channel.ID) (<-chan message.Message, <-chan error) {
	messages := make(chan message.Message)
	errs := make(chan error)

	go func() {
		defer close(messages)
		defer close(errs)

		iter := r.database.Query(_table.SelectAll()).WithContext(ctx).Iter()
		defer func() {
			err := iter.Close()
			if err != nil {
				errs <- err
				return
			}
		}()

		var model Model
		for iter.StructScan(&model) {
			msg, err := ToEntity(model)
			if err != nil {
				errs <- err
				return
			}

			messages <- msg
		}
	}()

	return messages, errs
}
