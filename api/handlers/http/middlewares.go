package http

import (
	"time"

	"github.com/rezamokaram/sample-ws/pkg/logger"

	appCtx "github.com/rezamokaram/sample-ws/pkg/context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func setUserContext(c *fiber.Ctx) error {
	c.SetUserContext(appCtx.NewAppContext(c.UserContext(), appCtx.WithLogger(logger.NewLogger())))
	return c.Next()
}

func ipRateLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        10,
		Expiration: time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests from this IP",
			})
		},
	})
}
