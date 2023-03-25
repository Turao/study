package users

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
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

func (UserRegistered) Process(data interface{}) interface{} {
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

	registered := event.Payload.Before == nil && event.Payload.After != nil
	if !registered {
		return nil
	}

	log.Println("user registered")
	return data // do nothing
}
