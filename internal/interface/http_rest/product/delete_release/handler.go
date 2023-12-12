package deleterelease

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// @Summary Releases reserve.
// @Description On input takes warehouse_id and product ids .
// @Tags 	warehouse
// @Accept	json
// @Produce	json
// @Param reqBody body requestBody true "input"
// @Success 200
// @Router 	/product/release [delete]
func (h *handler) Handle(c *fiber.Ctx) error {
	var req requestBody
	if err := c.BodyParser(&req); err != nil {
		return errors.Wrap(err, "failed to parse body")
	}

	if err := h.usecase.Execute(c.Context(), req.toUsecasePayload()); err != nil {
		return errors.Wrap(err, "failed to cancel reservations")
	}
	return c.SendStatus(fiber.StatusOK)
}
