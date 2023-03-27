package main

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"github.com/yusupovanton/moneyExchange/internal/me-scraper/app"
)

func handleError(err error) {
	log.Fatalf("A fatal error occured in the app module: %v", err)
}

func main() {

	ctx := context.Background()

	appConfig := app.NewAppConfig(ctx)
	appInstance := app.New(ctx, appConfig)

	appInstance.Run()
}
