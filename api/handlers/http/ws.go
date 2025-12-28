package http

import (
	"context"
	"encoding/json"

	"github.com/fasthttp/websocket"
	"github.com/rezamokaram/sample-ws/api/service"
	appCtx "github.com/rezamokaram/sample-ws/pkg/context"
	"github.com/valyala/fasthttp"

	"github.com/gofiber/fiber/v2"
)

var upgrader = websocket.FastHTTPUpgrader{
	CheckOrigin: func(ctx *fasthttp.RequestCtx) bool {
		return true
	},
}

func WS(svcGetter ServiceGetter[*service.StreamService]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return upgrader.Upgrade(ctx.Context(), func(conn *websocket.Conn) {
			wsContext := context.Background()
			svc := svcGetter(wsContext)

			stream, err := svc.GetStream(wsContext)
			if err != nil {
				appCtx.GetLogger(ctx.UserContext()).Error("get stream error:", "err", err)
				return
			}

			for msg := range stream.Channel {
				bytes, err := json.Marshal(msg)
				if err != nil {
					appCtx.GetLogger(ctx.UserContext()).Error("marshaling error:", "err", err)
					return
				}

				if err := conn.WriteMessage(websocket.TextMessage, bytes); err != nil {
					appCtx.GetLogger(ctx.UserContext()).Error("write error:", "err", err)
					return
				}
			}

			appCtx.GetLogger(ctx.UserContext()).Info("successful")
		})
	}
}
