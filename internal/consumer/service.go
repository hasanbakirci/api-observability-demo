package consumer

import (
	"fmt"

	"github.com/hasanbakirci/api-observability-demo/internal/models"
)

type ConsumerService struct {
}

func NewConsumerService() *ConsumerService {
	return &ConsumerService{}
}

func (c *ConsumerService) Process(event *models.Event) error {
	// Process the event
	fmt.Println("Processing event:", event)
	return nil
}
