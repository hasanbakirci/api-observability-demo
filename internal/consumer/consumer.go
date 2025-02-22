package consumer

import (
	"fmt"

	"github.com/hasanbakirci/api-observability-demo/internal/models"
	"github.com/hasanbakirci/api-observability-demo/pkg/rabbit"
)

type Consumer struct {
	service  ConsumerService
	consumer *rabbit.RabbitConsumer[models.Event]
}

func NewConsumer(service ConsumerService, consumer *rabbit.RabbitConsumer[models.Event]) *Consumer {
	return &Consumer{
		service:  service,
		consumer: consumer,
	}
}

func (c Consumer) Start() error {
	for {
		err := c.consumer.ConsumeMessage(func(e *models.Event) error {
			return c.service.Process(e)
		})
		if err != nil {
			fmt.Println("Error consuming message:", err)
		}
	}
}
