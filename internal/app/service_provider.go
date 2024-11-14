package app

import (
	"github.com/sergeyiksanov/notification-service/internal/config"
	"log"
)

type serviceProvider struct {
	rabbitMqConfig *config.RabbitMqConfig
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) getRabbitMqProvider() *config.RabbitMqConfig {
	if s.rabbitMqConfig == nil {
		r, err := config.NewRabbitMqConfig()

		if err != nil {
			log.Fatalf("Failed set rabbit mq config: %s", err)
		}

		s.rabbitMqConfig = r
	}

	return s.rabbitMqConfig
}
