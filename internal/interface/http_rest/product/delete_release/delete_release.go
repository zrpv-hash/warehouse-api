package deleterelease

import (
	"net/http"
	"warehousesvc/internal/application/reserve/release"
	"warehousesvc/internal/interface/http_rest/common"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase release.UseCase
}

func New(uc release.UseCase) common.Handler {
	return &handler{uc}
}

func (h *handler) Pattern() string {
	return "/product/release"
}

func (h *handler) Method() string {
	return http.MethodDelete
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
