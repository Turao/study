package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/turao/topics/streams/processor/users"
)

type Processor interface {
	Name() string
	Inbound() string
	Outbound() string

	Process(msg *message.Message) ([]*message.Message, error)
}

func main() {
	logger := watermill.NewStdLogger(false, false)

	router, err := message.NewRouter(
		message.RouterConfig{},
		logger,
	)
	if err != nil {
		log.Fatalln(err)
	}

	router.AddPlugin(plugin.SignalsHandler)

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:     []string{"localhost:9092"},
			Unmarshaler: kafka.DefaultMarshaler{},
		},
		logger,
	)
	if err != nil {
		log.Fatalln(err)
	}

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{"localhost:9092"},
			Marshaler: kafka.DefaultMarshaler{},
		},
		logger,
	)
	if err != nil {
		log.Fatalln(err)
	}

	processors := []Processor{
		users.UserRegistered{},
		users.EmailUpdated{},
		users.NameUpdated{},
	}

	// register handlers
	for _, processor := range processors {
		poisonMiddleware, err := middleware.PoisonQueue(
			publisher,
			fmt.Sprintf("%s.dlq", processor.Inbound()),
		)
		if err != nil {
			log.Println(err)
			continue // skip processor
		}

		handler := router.AddHandler(
			processor.Name(),
			processor.Inbound(),
			subscriber,
			processor.Outbound(),
			publisher,
			processor.Process,
		)

		handler.AddMiddleware(
			middleware.Retry{
				MaxRetries: 3,
			}.Middleware,
			poisonMiddleware,
			middleware.Recoverer,
		)
	}

	err = router.Run(context.Background())
	defer func() {
		err := router.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	if err != nil {
		log.Fatalln(err)
	}
}
