package config

type Config struct {
	HTTPPort             string
	RabbitBootsrapServer string
	RabbitExchangeName   string
	RabbitQueueName      string
	RabbitRoutingKey     string
}

func GetConfig() *Config {
	return &Config{
		HTTPPort:             "8080",
		RabbitBootsrapServer: "amqp://guest:guest@rabbitmq:5672/",
		RabbitExchangeName:   "EventExchange",
		RabbitQueueName:      "EventQueue",
		RabbitRoutingKey:     "EventQueue",
	}
}
