package reserve

import (
	"context"
	"warehousesvc/internal/infrastructure/tx/isolation"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) error {
	err := i.txManager.Do(ctx, isolation.LevelReadCommitted, func(ctx context.Context) error {
		inv, err := i.inventoryRepo.GetAllOptionsByWarehouse(ctx, p.WarehouseId, p.GetArray())
		if err != nil {
			return errors.Wrap(err, "failed to get options by warehouse")
		}

		for _, v := range inv {
			if p.UniqueCodeMap[v.OptionId] > v.Quantity {
				return errors.New("not enough products in a warehouse")
			}
		}

		for optionId, quantity := range p.UniqueCodeMap {
			if err = i.reserveRepo.Reserve(ctx, quantity, optionId, p.WarehouseId); err != nil {
				return errors.Wrap(err, "failed to reserve products")
			}
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to run transaction")
	}

	return nil
}
