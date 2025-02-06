package main

import (
	"context"
	"finance-operations-service/internal/app"
	"log"
	"os"
)

const portEnvName = "PORT"

func main() {
	ctx := context.Background()
	a := app.NewApp(ctx)
	err := a.Run(os.Getenv(portEnvName))
	if err != nil {
		log.Fatal(err)
	}
}
