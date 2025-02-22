package rabbit

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

type RabbitPublisher[T any] struct {
	connection    *amqp.Connection
	rabbitConnStr []string
	exchange      string
}

func NewRabbitPublisher[T any](rabbitConnStr []string, exchange string) (*RabbitPublisher[T], error) {
	conn, rabbitConnErr := connect(rabbitConnStr)
	if rabbitConnErr != nil {
		return nil, rabbitConnErr
	}

	return &RabbitPublisher[T]{
		connection:    conn,
		rabbitConnStr: rabbitConnStr,
		exchange:      exchange,
	}, nil
}

func (r *RabbitPublisher[T]) PublishMessage(message *T, routingKey string) error {
	ch, rabbitChErr := r.connection.Channel()
	if rabbitChErr != nil {
		return rabbitChErr
	}
	defer ch.Close()
	msg, jsonErr := json.Marshal(message)
	if jsonErr != nil {
		return jsonErr
	}
	return ch.Publish(
		r.exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		})
}
