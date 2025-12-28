package context

import (
	"context"
	"log/slog"
	"os"
)

var defaultLogger *slog.Logger

func init() {
	defaultLogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

type appContext struct {
	context.Context
	shouldCommit bool
	logger       *slog.Logger
}

type AppContextOpt func(*appContext) *appContext // option pattern

func WithLogger(logger *slog.Logger) AppContextOpt {
	return func(ac *appContext) *appContext {
		ac.logger = logger
		return ac
	}
}

func NewAppContext(parent context.Context, opts ...AppContextOpt) context.Context {
	ctx := &appContext{Context: parent}
	for _, opt := range opts {
		ctx = opt(ctx)
	}

	return ctx
}

func SetLogger(ctx context.Context, logger *slog.Logger) {
	if appCtx, ok := ctx.(*appContext); ok {
		appCtx.logger = logger
	}
}

func GetLogger(ctx context.Context) *slog.Logger {
	appCtx, ok := ctx.(*appContext)
	if !ok || appCtx.logger == nil {
		return defaultLogger
	}

	return appCtx.logger
}
