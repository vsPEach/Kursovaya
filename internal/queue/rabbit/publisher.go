package rabbit

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"log"
)

type Publisher struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	logg       zap.Logger
}

func NewPublisher() (*Publisher, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &Publisher{connection: conn, channel: channel}, nil
}

func (p *Publisher) Publish(ctx context.Context, body []byte) {
	if p.connection == nil || p.connection.IsClosed() {
		log.Println("")
	}

	queue, err := p.channel.QueueDeclare("queue1", false, false, false, false, nil)
	if err != nil {
		log.Println(err)
	}

	err = p.channel.PublishWithContext(ctx, "", queue.Name, false, false, amqp.Publishing{
		ContentType:  "application/json",
		Body:         body,
		DeliveryMode: amqp.Persistent,
	})

	if err != nil {
		log.Println(err)
	}
}
