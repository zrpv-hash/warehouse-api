package product

import "context"

type Repository interface {
	GetAllProductsFromWarehouse(context.Context, string, int) ([]Product, error)
}
