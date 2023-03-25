package main

import (
	"context"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/reugn/go-streams/flow"
	gostreams "github.com/reugn/go-streams/kafka"
	"github.com/turao/topics/streams/processor/users"
)

type Processor interface {
	Name() string
	Inbound() string
	Outbound() string

	Process(event interface{}) interface{}
}

func main() {
	addresses := []string{"localhost:9092"}

	saramacfg := sarama.NewConfig()
	saramacfg.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	saramacfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	saramacfg.Producer.Return.Successes = true

	processors := []Processor{
		users.UserRegistered{},
		users.EmailUpdated{},
		users.NameUpdated{},
	}

	wg := sync.WaitGroup{}
	for _, processor := range processors {
		wg.Add(1)
		go func(processor Processor) {
			defer wg.Done()
			source := gostreams.NewKafkaSource(
				context.Background(),
				addresses,
				processor.Name(),
				saramacfg,
				processor.Inbound(),
			)

			sink := gostreams.NewKafkaSink(
				addresses,
				saramacfg,
				processor.Outbound(),
			)

			mapper := flow.NewMap(processor.Process, 1)

			source.
				Via(mapper).
				To(sink)
		}(processor)
	}
	wg.Wait()
}
