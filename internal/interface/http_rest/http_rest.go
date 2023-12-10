package httprest

import (
	"context"
	"fmt"
	"time"
	"warehousesvc/internal/core/config"
	"warehousesvc/internal/interface/http_rest/common"

	"github.com/gofiber/fiber/v2"
)

func New(ctx context.Context, config *config.Config, handlers []common.Handler) {
	f := fiber.New(
		fiber.Config{
			ReadTimeout: time.Second * 3,
		},
	)

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
	f.Shutdown()
}
