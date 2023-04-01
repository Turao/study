package message

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/turao/topics/messages/entity/message"
)

type repository struct {
	messages map[string]*Model
	mutex    sync.RWMutex
}

func NewRepository() (*repository, error) {
	return &repository{
		messages: make(map[string]*Model, 0),
		mutex:    sync.RWMutex{},
	}, nil
}

func (r *repository) Save(ctx context.Context, message message.Message) error {
	msg, err := ToModel(message)
	if err != nil {
		return err
	}

	r.mutex.Lock()
	r.messages[msg.ID] = msg
	r.mutex.Unlock()

	r.printMessages()
	return nil
}

func (r *repository) printMessages() {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	pretty, _ := json.MarshalIndent(r.messages, "", " ")
	log.Println("message stored:", string(pretty))
}
