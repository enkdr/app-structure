package main

import (
	"context"
	"fmt"
	"app-structure/application"
	"os"
	"os/signal"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	err := app.Start(ctx)

	if err != nil {
		fmt.Println("failed to start app", err)
	}

}
