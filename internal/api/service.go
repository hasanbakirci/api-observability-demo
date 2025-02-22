package api

import (
	"fmt"

	"github.com/hasanbakirci/api-observability-demo/config"
	"github.com/hasanbakirci/api-observability-demo/internal/models"
	"github.com/hasanbakirci/api-observability-demo/pkg/rabbit"
)

type Service struct {
	config    *config.Config
	publisher *rabbit.RabbitPublisher[models.Event]
}

func NewService(config *config.Config, publisher *rabbit.RabbitPublisher[models.Event]) *Service {
	return &Service{
		config:    config,
		publisher: publisher,
	}
}

func (s *Service) PublishEvent(event models.Event) error {
	err := s.publisher.PublishMessage(&event, s.config.RabbitRoutingKey)
	if err != nil {
		fmt.Println("Error publishing message:", err)
		return err
	}
	return nil
}
