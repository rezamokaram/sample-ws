package http

import (
	"context"
	"fmt"

	"github.com/rezamokaram/sample-ws/api/service"
	"github.com/rezamokaram/sample-ws/app"
	"github.com/rezamokaram/sample-ws/config"
)

// stream service transient instance handler
func streamServiceGetter(appContainer app.App, cfg config.ServerConfig) ServiceGetter[*service.StreamService] {
	return func(ctx context.Context) *service.StreamService {
		fmt.Printf("PLACE-3 \n\n\n")
		return service.NewStreamService(appContainer.ProducerService(ctx))
	}
}
