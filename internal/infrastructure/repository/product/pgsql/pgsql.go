package pgsql

import (
	"context"
	"database/sql"
	"warehousesvc/internal/domain/product"
	"warehousesvc/internal/infrastructure/tx/pgsqltx"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) product.Repository {
	return &repo{db}
}

const (
	limit int = 50
)

func (r *repo) GetAllProductsFromWarehouse(ctx context.Context, id string, offset int) ([]product.Product, error) {
	const getAllByWarehouseIdQuery = `
		SELECT
			option.id AS id,
			product.name AS name,
			size.name AS size,
			inventory.quantity AS quantity
		FROM
			product
		JOIN option ON product.id = option.product_id
		JOIN size ON option.size = size.name
		JOIN inventory ON option.id = inventory.option_id
		WHERE
			inventory.warehouse_id = $1
		ORDER BY
			id
		LIMIT $2
		OFFSET $3;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row []productRow

	err := q.SelectContext(ctx, &row, getAllByWarehouseIdQuery, id, limit, limit*offset)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return toDomainArray(row), nil
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return toDomainArray(row), nil
}
