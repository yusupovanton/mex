package app

import (
	"context"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	DBConfig *DBConfig
}

type DBConfig struct {
	DSN string `envconfig:"DB_DSN_POSTGRES"`
}

func NewAppConfig(ctx context.Context) *AppConfig {

	var err error
	var dbconfig = DBConfig{}

	log.Println(os.Getenv("DB_DSN_POSTGRES"))

	err = envconfig.Process("", &dbconfig)

	if err != nil {
		log.Fatalf("Failed to process envs: %v", err)
	}

	return &AppConfig{DBConfig: &dbconfig}
}
