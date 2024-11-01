package app

import (
	"context"
	"github.com/sergeyiksanov/notification-service/internal/config"
	"github.com/sergeyiksanov/notification-service/internal/handler"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	serviceProvider *serviceProvider
	rabbitMqChannel *amqp091.Channel
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) {
	err := a.initDeps(ctx)
	if err != nil {
		log.Fatalf("failed to init deps: %s", err)
	}

	err = a.connectRabbitMq(ctx)
	if err != nil {
		log.Fatalf("failed to connect rabbit mq: %s", err)
	}
	defer a.rabbitMqChannel.Close()

	a.startConsumers(ctx)

	select {
	case <-ctx.Done():
		return
	}

}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.connectRabbitMq,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	return err
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) connectRabbitMq(_ context.Context) error {
	conn, err := amqp091.Dial(a.serviceProvider.getRabbitMqProvider().URL)
	if err != nil {
		log.Fatalf("Failed to connect to rabbit mq: %s", err)
	}
	log.Print("Connect to rabbit mq")

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel rabbit mq: %s", err)
	}
	log.Print("Open a channel rabbit mq")

	for _, queueName := range a.serviceProvider.getRabbitMqProvider().Queues {
		_, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
		if err != nil {
			log.Fatalf("failed to declare queue: %s", err)
		}
		log.Printf("Declate queue: %s", queueName)
	}

	a.rabbitMqChannel = channel
	log.Print("Set rabbit mq channel")

	return nil
}

func (a *App) startConsumers(ctx context.Context) {
	messagesEvent, err := a.rabbitMqChannel.Consume("event_notifications", "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("failed to register consumer: %s", err)
	}

	messagesBroadcast, err := a.rabbitMqChannel.Consume("broadcast_notifications", "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("failed to register consumer: %s", err)
	}

	go func() {
		for {
			select {
			case msg := <-messagesEvent:
				h := handler.EventNotificationHandler{}
				h.Handle(msg)
			case <-ctx.Done():
				log.Print("Shuting down consumer...")
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case msg := <-messagesBroadcast:
				log.Printf("Received message: %s", msg.Body)
			case <-ctx.Done():
				return
			}
		}
	}()
}
