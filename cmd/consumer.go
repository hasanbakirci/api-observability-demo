package cmd

import (
	"fmt"

	"github.com/hasanbakirci/api-observability-demo/config"
	"github.com/hasanbakirci/api-observability-demo/internal/consumer"
	"github.com/hasanbakirci/api-observability-demo/internal/models"
	"github.com/hasanbakirci/api-observability-demo/pkg/rabbit"
)

type ConsumerCmd struct {
	instance consumer.Consumer
	config   *config.Config
}

func RunConsumer() {
	fmt.Println("Running Consumer")
	var consumerCmd = &ConsumerCmd{
		config: config.GetConfig(),
	}

	rabbitConsumer, rabbitConsumerErr := rabbit.NewRabbitConsumer[models.Event](
		[]string{consumerCmd.config.RabbitBootsrapServer},
		consumerCmd.config.RabbitExchangeName,
		consumerCmd.config.RabbitQueueName,
		consumerCmd.config.RabbitRoutingKey,
	)
	if rabbitConsumerErr != nil {
		panic(rabbitConsumerErr)
	}

	service := consumer.NewConsumerService()
	consumerCmd.instance = *consumer.NewConsumer(*service, rabbitConsumer)
	if err := consumerCmd.instance.Start(); err != nil {
		fmt.Println("Error starting Consumer:", err)
	}
}
