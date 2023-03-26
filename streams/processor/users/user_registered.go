package users

import (
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gofrs/uuid"

	eventsV1 "github.com/turao/topics/events/users/v1"
)

type UserRegistered struct{}

func (UserRegistered) Name() string {
	return "registered"
}

func (UserRegistered) Inbound() string {
	return "cdc.public.users"
}

func (UserRegistered) Outbound() string {
	return "user.registered"
}

func (UserRegistered) Process(msg *message.Message) ([]*message.Message, error) {
	var event CDCEvent
	err := json.Unmarshal(msg.Payload, &event)
	if err != nil {
		log.Println("failed to unmarshal message")
		return nil, err
	}

	registered := event.Payload.Before == nil && event.Payload.After != nil
	if !registered {
		return nil, nil
	}

	log.Println("user registered")
	evt := eventsV1.UserRegistered{
		ID:        event.Payload.After.ID,
		Email:     event.Payload.After.Email,
		FirstName: event.Payload.After.Firstname,
		LastName:  event.Payload.After.Lastname,
	}

	payload, err := json.Marshal(evt)
	if err != nil {
		return nil, err
	}

	return []*message.Message{
		message.NewMessage(
			uuid.Must(uuid.NewV4()).String(),
			payload,
		),
	}, nil
}
