package cmd

import (
	"fmt"

	"github.com/hasanbakirci/api-observability-demo/config"
	"github.com/hasanbakirci/api-observability-demo/internal/api"
	"github.com/hasanbakirci/api-observability-demo/internal/models"
	"github.com/hasanbakirci/api-observability-demo/pkg/rabbit"
	"github.com/labstack/echo/v4"
)

type ApiCmd struct {
	instance *echo.Echo
	config   *config.Config
}

func RunApi() {
	fmt.Println("Running API")
	var apiCmd = &ApiCmd{
		instance: echo.New(),
		config:   config.GetConfig(),
	}

	publisher, publisherErr := rabbit.NewRabbitPublisher[models.Event](
		[]string{apiCmd.config.RabbitBootsrapServer},
		apiCmd.config.RabbitExchangeName,
	)
	if publisherErr != nil {
		panic(publisherErr)
	}
	service := api.NewService(apiCmd.config, publisher)
	api.NewHandler(apiCmd.instance, service)

	if err := apiCmd.instance.Start(":" + apiCmd.config.HTTPPort); err != nil {
		fmt.Println("Error starting API:", err)
	}
}
