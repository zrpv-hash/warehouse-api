package postreserve

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// @Summary Reserves products if warehouse have enough.
// @Description On input takes warehouse_id, product ids and  quantities.
// @Tags 	warehouse
// @Accept	json
// @Produce	json
// @Param reqBody body requestBody true "input"
// @Success 200
// @Router 	/product/reserve [post]
func (h *handler) Handle(c *fiber.Ctx) error {
	var body requestBody
	if err := c.BodyParser(&body); err != nil {
		return errors.Wrap(err, "failed to parse body")
	}
	err := h.usecase.Execute(c.Context(), body.toUsecasePayload())
	if err != nil {
		return errors.Wrap(err, "failed to reserve products")
	}
	return c.SendStatus(http.StatusCreated)
}
