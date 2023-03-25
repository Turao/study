package main

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/reugn/go-streams/flow"
	gostreams "github.com/reugn/go-streams/kafka"
	"github.com/turao/topics/streams/processor/users"
)

type Processor interface {
	Inbound() string
	Outbound() string

	Process(event interface{}) interface{}
}

func main() {
	addresses := []string{"localhost:9092"}
	groupID := "streams-2"

	saramacfg := sarama.NewConfig()
	saramacfg.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	saramacfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	saramacfg.Producer.Return.Successes = true

	processors := []Processor{
		users.EmailUpdated{},
	}

	for _, processor := range processors {
		source := gostreams.NewKafkaSource(
			context.Background(),
			addresses,
			groupID,
			saramacfg,
			processor.Inbound(),
		)

		sink := gostreams.NewKafkaSink(
			addresses,
			saramacfg,
			processor.Outbound(),
		)

		processor := flow.NewMap(processor.Process, 1)

		source.
			Via(processor).
			To(sink)
	}
}
