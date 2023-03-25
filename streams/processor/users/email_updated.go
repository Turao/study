package users

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
)

type EmailUpdated struct{}

func (EmailUpdated) Inbound() string {
	return "cdc.public.users"
}

func (EmailUpdated) Outbound() string {
	return "user.email-updated"
}

func (EmailUpdated) Process(data interface{}) interface{} {
	log.Println("transforming message")

	msg, ok := data.(*sarama.ConsumerMessage)
	if !ok {
		return nil
	}

	if msg == nil {
		return nil
	}

	var event CDCEvent
	err := json.Unmarshal(msg.Value, &event)
	if err != nil {
		log.Println("failed to unmarshal message")
		return nil
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
		return nil
	}

	log.Println("user email updated")
	return data // do nothing
}
