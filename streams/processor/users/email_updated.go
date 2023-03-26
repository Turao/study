package users

import (
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
)

type EmailUpdated struct{}

func (EmailUpdated) Name() string {
	return "email-updated"
}

func (EmailUpdated) Inbound() string {
	return "cdc.public.users"
}

func (EmailUpdated) Outbound() string {
	return "user.email-updated"
}

func (EmailUpdated) Process(msg *message.Message) ([]*message.Message, error) {
	var event CDCEvent
	err := json.Unmarshal(msg.Payload, &event)
	if err != nil {
		log.Println("failed to unmarshal message")
		return nil, err
	}

	// check if email has changed
	var before string
	if event.Payload.Before != nil {
		before = event.Payload.Before.Email
	}

	var after string
	if event.Payload.After != nil {
		after = event.Payload.After.Email
	}

	if before == after {
		return nil, nil
	}

	log.Println("user email updated")
	// do nothing
	return []*message.Message{
		msg,
	}, nil
}
