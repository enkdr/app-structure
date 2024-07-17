package main

import (
	"app-structure/application"
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	app := application.NewApp()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	err := app.Start(ctx)

	if err != nil {
		fmt.Println("failed to start app", err)
	}

}
