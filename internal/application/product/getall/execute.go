package getall

import (
	"context"
	"warehousesvc/internal/domain/product"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) ([]product.Product, error) {
	res, err := i.productRepo.GetAllProductsFromWarehouse(ctx, p.WarehouseId, p.Page)
	if err != nil {
		errors.Wrap(err, "can't get all products from warehouse")
	}

	return res, nil
}
