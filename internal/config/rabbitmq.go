package config

import (
	"errors"
	"os"
)

type RabbitMqConfig struct {
	URL    string
	Queues map[string]string
}

func NewRabbitMqConfig() (*RabbitMqConfig, error) {
	url := os.Getenv("RABBITMQ_URL")
	if len(url) == 0 {
		return nil, errors.New("environment variable RABBITMQ_URL not initialized")
	}

	queueEventNotifications := os.Getenv("RABBITMQ_QUEUE_EVENT_NOTIFICATIONS")
	if len(queueEventNotifications) == 0 {
		return nil, errors.New("environment variable RABBITMQ_QUEUE_EVENT_NOTIFICATIONS not initialized")
	}

	queueBroadcastNotifications := os.Getenv("RABBITMQ_QUEUE_BROADCAST_NOTIFICATIONS")
	if len(queueBroadcastNotifications) == 0 {
		return nil, errors.New("environment variable RABBITMQ_QUEUE_EVENT_NOTIFICATIONS not initialized")
	}

	return &RabbitMqConfig{
		URL: url,
		Queues: map[string]string{
			"event_notifications":     queueEventNotifications,
			"broadcast_notifications": queueBroadcastNotifications,
		},
	}, nil
}
