package app

import (
	"context"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"

	"github.com/yusupovanton/moneyExchange/internal/me-scraper/service"
)

type application struct {
	db             *sqlx.DB
	config         *AppConfig
	service        *service.ScraperService
	backgroundJobs []func() error
	stopHandlers   []func()
	ctx            context.Context
}

type Application interface {
	RegisterBackgroundJob(backgroundJob func() error)
	RegisterStopHandler(stopHandler func())
	Run() error
}

func New(ctx context.Context, config *AppConfig) *application {

	newApp := application{}

	newApp.ctx = ctx
	newApp.config = config

	newApp.initDB(newApp.config.DBConfig)

	newApp.service = service.NewScraperService(newApp.db, newApp.ctx)
	newApp.initParser()
	return &newApp
}

func (app *application) initParser() {

	app.RegisterBackgroundJob(func() error {
		for {
			if err := app.service.ScrapeBinance(); err != nil {
				return err
			}
		}
	})

}

func (app *application) initDB(cfg *DBConfig) error {

	var err error

	app.db, err = sqlx.Open("postgres", cfg.DSN)
	log.Println(cfg.DSN)
	if err != nil {
		log.Fatalf("There was an error connecting to db: %v", err)
		return err
	}

	err = app.db.Ping()
	if err != nil {
		log.Fatalf("Could not ping db: %v", err)
		return err
	}
	app.RegisterStopHandler(func() {
		err = app.db.Close()
		if err != nil {
			log.Printf("Unable close db connection: %v", err)
		}
	})

	return nil

}

func (app *application) Run() error {
	log.Printf("Starting app")
	defer app.stop()
	errors := app.startBackgroundJobs()

	select {
	case <-app.ctx.Done():
		return nil
	case err := <-errors:
		return err
	}
}

func (app *application) RegisterBackgroundJob(backgroundJob func() error) {
	app.backgroundJobs = append(app.backgroundJobs, backgroundJob)
}

func (app *application) RegisterStopHandler(stopHandler func()) {
	app.stopHandlers = append(app.stopHandlers, stopHandler)
}

func (app *application) startBackgroundJobs() chan error {
	errors := make(chan error)

	for _, job := range app.backgroundJobs {
		_job := job // to prevent variable override during loop iterations
		go func() {
			errors <- _job()
		}()
	}

	return errors
}

func (app *application) stop() {
	var wg sync.WaitGroup
	wg.Add(len(app.stopHandlers))
	for _, handler := range app.stopHandlers {
		_handler := handler // to prevent variable override during loop iterations
		go func() {
			defer wg.Done()
			_handler()
		}()
	}
	wg.Wait()
}
