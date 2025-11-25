package sse

import (
	"fmt"
)

type KeepAliveEvent []byte

const _keepAliveEventContent string = ":keep-alive\n"

func (e KeepAliveEvent) Bytes() []byte {
	return []byte(_keepAliveEventContent)
}

type Event interface {
	Bytes() []byte
}

type DataEvent struct {
	Event string
	Data  []byte
	ID    *string
	Retry *int
}

func (e DataEvent) Bytes() []byte {
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
	return []byte(message)
}
