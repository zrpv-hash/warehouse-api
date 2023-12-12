package pgsql

import "warehousesvc/internal/domain/product"

type productRow struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Size     string `db:"size"`
	Quantity int    `db:"quantity"`
}

func (r *productRow) ToDomain() *product.Product {
	return product.FormData(r.ID, r.Name, r.Size, r.Quantity)
}

func toDomainArray(pr []productRow) []product.Product {
	res := make([]product.Product, len(pr))

	for k, v := range pr {
		res[k] = *v.ToDomain()
	}

	return res
}
