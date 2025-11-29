package user

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/turao/topics/users/entity/user"
)

const (
	topicName = "cdc.public.users"
)

// repository is the implementation of the user repository
type repository struct {
	subscriber message.Subscriber
}

// NewRepository creates a new user stream repository
func NewRepository(subscriber message.Subscriber) (*repository, error) {
	if subscriber == nil {
		return nil, errors.New("message subscriber is nil")
	}

	return &repository{
		subscriber: subscriber,
	}, nil
}

func (r *repository) StreamUsers(ctx context.Context) (<-chan user.User, error) {
	users := make(chan user.User)

	messages, err := r.subscriber.Subscribe(ctx, topicName)
	if err != nil {
		close(users)
		return users, err
	}

	go func() error {
		defer close(users)

		for message := range messages {
			user, err := fromMessageToAfterUser(message)
			if err != nil {
				return err
			}
			if user == nil {
				continue // skip
			}
			users <- user
			message.Ack()
		}

		return nil
	}()

	return users, nil
}

func fromMessageToAfterUser(message *message.Message) (user.User, error) {
	var event Event
	err := json.Unmarshal(message.Payload, &event)
	if err != nil {
		return nil, err
	}

	model := event.Payload.After
	if model == nil {
		return nil, nil
	}

	return ToEntity(*model)
}
