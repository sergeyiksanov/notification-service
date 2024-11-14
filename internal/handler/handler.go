package handler

import (
	req "github.com/sergeyiksanov/notification-service/pkg/api/v1"
	"log"

	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

type EventNotificationHandler struct{}

func (h *EventNotificationHandler) Handle(msg amqp091.Delivery) error {
	var event req.EventNotificationRequest
	err := proto.Unmarshal(msg.Body, &event)
	if err != nil {
		return err
	}

	log.Printf("Received message: %s", &event)
	return nil
}
