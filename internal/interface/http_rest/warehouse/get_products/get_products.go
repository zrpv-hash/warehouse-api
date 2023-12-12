package getproducts

import (
	"net/http"
	"warehousesvc/internal/application/product/getall"
	"warehousesvc/internal/interface/http_rest/common"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	getall.UseCase
}

func New(uc getall.UseCase) common.Handler {
	return &handler{uc}
}

func (h *handler) Pattern() string {
	return "/warehouse/product"
}

func (h *handler) Method() string {
	return http.MethodGet
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
