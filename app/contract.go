package app

import (
	"context"

	"github.com/rezamokaram/sample-ws/config"
	producerPort "github.com/rezamokaram/sample-ws/internal/producer/port"
)

type App interface {
	ProducerService(ctx context.Context) producerPort.Service
	Config() config.Config
}
