package getproducts

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// @Summary 	Returns all products from specified warehouse.
// @Description On input takes warehouse_id, and page offset.
// @Tags 		warehouse
// @Produce		json
// @Param      	id    	query	string  true  "id"
// @Param      	page    query   int  	true  "page offset"
// @Success 	200 {object} 	responseBody
// @Router 		/warehouse/product [get]
func (h *handler) Handle(c *fiber.Ctx) error {
	var rp reqParams
	if err := c.QueryParser(&rp); err != nil {
		return errors.Wrap(err, "failed to parse query")
	}

	r, err := h.UseCase.Execute(c.Context(), rp.toUsecasePayload())
	if err != nil {
		return errors.Wrap(err, "failed to get products")
	}

	resp := responseFromResult(&r)
	return c.JSON(resp)
}
