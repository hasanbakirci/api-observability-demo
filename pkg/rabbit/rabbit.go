package rabbit

import (
	"github.com/streadway/amqp"
)

func connect(rabbitConnStr []string) (*amqp.Connection, error) {
	conn := &amqp.Connection{}
	var rabbitConnErr error
	for _, s := range rabbitConnStr {
		conn, rabbitConnErr = amqp.Dial(s)
	}
	if rabbitConnErr != nil {
		return nil, rabbitConnErr
	}
	return conn, rabbitConnErr
}
