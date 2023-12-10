package deleterelease

import (
	"net/http"
	"warehousesvc/internal/interface/http_rest/common"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
}

func New() common.Handler {
	return &handler{}
}

func (h *handler) Pattern() string {
	return "/product/release"
}

func (h *handler) Method() string {
	return http.MethodGet
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
