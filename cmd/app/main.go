package main

import (
	"NotificationService/internal/app"
	"context"
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
