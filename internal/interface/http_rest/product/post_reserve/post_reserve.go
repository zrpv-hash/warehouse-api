package postreserve

import (
	"net/http"
	"warehousesvc/internal/application/reserve/reserve"
	"warehousesvc/internal/interface/http_rest/common"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase reserve.UseCase
}

func New(uc reserve.UseCase) common.Handler {
	return &handler{uc}
}

func (h *handler) Pattern() string {
	return "/product/reserve"
}

func (h *handler) Method() string {
	return http.MethodPost
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
