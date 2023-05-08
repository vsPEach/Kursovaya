package rabbit

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Subscriber struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewSubscriber() (*Subscriber, error) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}
	return &Subscriber{
		connection: connection,
		channel:    channel,
	}, nil
}

func (s *Subscriber) Subscribe(ctx context.Context) error {
	msgs, err := s.channel.Consume("queue1", "", false, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		for msg := range msgs {
			log.Println(msg.Body)
		}
	}()
	<-ctx.Done()
	return nil
}
