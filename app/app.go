package app

import (
	"context"

	"github.com/rezamokaram/sample-ws/config"
	"github.com/rezamokaram/sample-ws/internal/producer"
	producerPort "github.com/rezamokaram/sample-ws/internal/producer/port"
)

type app struct {
	cfg             config.SampleAuthConfig
	producerService producerPort.Service
}

func (a *app) ProducerService(ctx context.Context) producerPort.Service {
	return a.producerService
}

func (a *app) Config() config.Config {
	return a.cfg
}

func NewApp(cfg config.SampleAuthConfig) (App, error) {
	a := &app{
		cfg:             cfg,
		producerService: producer.NewService(),
	}
	return a, nil
}

func NewMustApp(cfg config.SampleAuthConfig) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
