package pgsql

import (
	"warehousesvc/internal/domain/inventory"
)

type inventoryRow struct {
	OptionId string `db:"option_id"`
	Quantity int    `db:"quantity"`
}

func (r *inventoryRow) ToDomain() *inventory.Inventory {
	return inventory.FormData(r.OptionId, r.Quantity)
}

func toDomainArray(ir []inventoryRow) []inventory.Inventory {
	res := make([]inventory.Inventory, len(ir))

	for _, v := range ir {
		res = append(res, *v.ToDomain())
	}

	return res
}
