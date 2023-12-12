package httprest

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	_ "warehousesvc/docs"
	"warehousesvc/internal/core/config"
	"warehousesvc/internal/interface/http_rest/common"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func New(ctx context.Context, config *config.Config, log *slog.Logger, handlers []common.Handler) {
	f := fiber.New(
		fiber.Config{
			ReadTimeout: time.Second * 3,
		},
	)

	f.Use(logger.New())

	f.Get("/swagger/*", swagger.HandlerDefault)

	for _, h := range handlers {
		f.Add(h.Method(), h.Pattern(), append(h.Middleware(), h.Handle)...)
	}

	go func() {
		err := f.Listen(fmt.Sprintf("%s:%d", config.HttpServer.Host, config.HttpServer.Port))
		if err != nil {
			panic(err)
		}
	}()

	<-ctx.Done()
	log.Info("Gracefully stopping Fiber Server")
	f.Shutdown()
}
