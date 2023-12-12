package release

import (
	"context"
	"warehousesvc/internal/infrastructure/tx/isolation"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) error {
	err := i.txManager.Do(ctx, isolation.LevelReadCommitted, func(ctx context.Context) error {
		if err := i.inventoryRepo.UpdateInventory(ctx, p.WarehouseId, p.GetArray()); err != nil {
			return errors.Wrap(err, "failed to update inventory")
		}

		if err := i.reserveRepo.DeleteReserved(ctx, p.GetArray(), p.WarehouseId); err != nil {
			return errors.Wrap(err, "failed to delete rows from reserved_products")
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to run transaction")
	}

	return nil
}
