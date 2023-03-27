package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/yusupovanton/moneyExchange/internal/me-scraper/app"
)

func handleError(err error) {
	if err != nil {
		log.Fatalf("A fatal error occured in the app module: %v", err)
	}

}

func main() {

	err := godotenv.Load("cmd/.env")
	handleError(err)

	alt := os.Getenv("DB_DSN_POSTGRES")
	log.Println(alt)
	if alt == "" {
		log.Fatalln(alt)
	}

	ctx := context.Background()

	appConfig := app.NewAppConfig(ctx)

	appInstance := app.New(ctx, appConfig)

	appInstance.Run()
}
