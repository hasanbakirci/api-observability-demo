package rabbit

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

type RabbitConsumer[T any] struct {
	connection    *amqp.Connection
	rabbitConnStr []string
	exchange      string
	queue         string
	routingKey    string
}

func NewRabbitConsumer[T any](rabbitConnStr []string, exchange, queue, routingKey string) (*RabbitConsumer[T], error) {
	conn, rabbitConnErr := connect(rabbitConnStr)
	if rabbitConnErr != nil {
		return nil, rabbitConnErr
	}
	return &RabbitConsumer[T]{
		connection:    conn,
		rabbitConnStr: rabbitConnStr,
		exchange:      exchange,
		queue:         queue,
		routingKey:    routingKey,
	}, nil
}

func (r *RabbitConsumer[T]) ConsumeMessage(handler func(*T) error) error {
	ch, err := r.connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		r.queue,
		"",
		false, // autoAck: false
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for d := range msgs {
		var message T
		if err := json.Unmarshal(d.Body, &message); err != nil {
			return err
		}
		if err := handler(&message); err != nil {
			d.Ack(false)
			return err
		}
		d.Ack(false)
	}
	return nil
}
