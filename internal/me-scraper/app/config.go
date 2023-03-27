package app

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
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

	err = envconfig.Process(ctx, &dbconfig)

	if err != nil {
		log.Fatalf("Failed to process envs: %v", err)
	}

	return &AppConfig{DBConfig: &dbconfig}
}
