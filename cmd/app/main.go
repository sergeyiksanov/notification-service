package main

import (
	"context"
	"github.com/sergeyiksanov/notification-service/internal/app"
	"log"
)

func main() {
	ctx, cansel := context.WithCancel(context.Background())
	defer cansel()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to create app: %s", err)
	}

	a.Run(ctx)
}
