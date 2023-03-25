package users

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

type NameUpdated struct{}

func (NameUpdated) Name() string {
	return "name-updated"
}

func (NameUpdated) Inbound() string {
	return "cdc.public.users"
}

func (NameUpdated) Outbound() string {
	return "user.name-updated"
}

func (NameUpdated) Process(data interface{}) interface{} {
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
		before = formatName(
			event.Payload.Before.Firstname,
			event.Payload.Before.Lastname,
		)
	}

	var after string
	if event.Payload.After != nil {
		after = formatName(
			event.Payload.After.Firstname,
			event.Payload.After.Lastname,
		)
	}

	if before == after {
		return nil
	}

	log.Println("user name updated")
	return data // do nothing
}

func formatName(firstname string, lastname string) string {
	return fmt.Sprintf("%s %s", firstname, lastname)
}
