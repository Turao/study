package sse

import (
	"fmt"
)

const KeepAlive string = ":keep-alive\n"

type Event struct {
	Event string
	Data  []byte
	ID    *string
	Retry *int
}

func (e Event) String() string {
	var message string
	message += fmt.Sprintf("name: %s\n", e.Event)
	message += fmt.Sprintf("data: %s\n", e.Data)
	if e.ID != nil {
		message += fmt.Sprintf("id: %s\n", *e.ID)
	}
	if e.Retry != nil {
		message += fmt.Sprintf("retry: %v\n", *e.Retry)
	}
	message += "\n"
	return message
}
