package http

import (
	"fmt"

	"github.com/rezamokaram/sample-ws/app"
	"github.com/rezamokaram/sample-ws/config"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	router := fiber.New()

	router.Use(ipRateLimiter())

	api := router.Group("/api/v1", setUserContext)

	registerStreamAPI(appContainer, cfg, api)

	return router.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}

func registerStreamAPI(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	userSvcGetter := streamServiceGetter(appContainer, cfg)
	router.Get("/ws", WS(userSvcGetter))
}
