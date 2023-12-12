package inventory

import "context"

type Repository interface {
	GetAllOptionsByWarehouse(context.Context, string, []string) ([]Inventory, error)
	UpdateInventory(context.Context, string, []string) error
}
