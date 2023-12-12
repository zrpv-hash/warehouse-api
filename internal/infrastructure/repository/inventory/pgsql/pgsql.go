package pgsql

import (
	"context"
	"database/sql"
	"warehousesvc/internal/domain/inventory"
	"warehousesvc/internal/domain/product"
	"warehousesvc/internal/infrastructure/tx/pgsqltx"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) inventory.Repository {
	return &repo{db}
}

func (r *repo) GetAllOptionsByWarehouse(ctx context.Context, warehouseId string, optionIds []string) ([]inventory.Inventory, error) {
	const getAllOptionsByWarehouseQuery = `
		SELECT
			i.option_id AS option_id,
			i.quantity AS quantity
		FROM
			inventory i
		WHERE
			i.warehouse_id = $1
			AND i.option_id = ANY($2);
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row []inventoryRow

	err := q.SelectContext(ctx, &row, getAllOptionsByWarehouseQuery, warehouseId, pq.Array(optionIds))
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

func (r *repo) UpdateInventory(ctx context.Context, warehouseId string, optionIds []string) error {
	const updateInventory = `
		UPDATE inventory
		SET quantity = quantity + pr.total_quantity_reserved
		FROM (
			SELECT
				pr.unique_code_id,
				pr.warehouse_id,
				SUM(pr.quantity_reserved) AS total_quantity_reserved
			FROM product_reservation pr
			WHERE pr.warehouse_id = $1
			AND pr.unique_code_id = ANY($2)
			GROUP BY pr.unique_code_id, pr.warehouse_id
		) pr
		WHERE inventory.option_id = pr.unique_code_id
		AND inventory.warehouse_id = pr.warehouse_id;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)

	_, err := q.ExecContext(ctx, updateInventory, warehouseId, pq.Array(optionIds))
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return errors.Wrap(product.ErrProductNotFound, "products in specified warehouse not found")
		default:
			return errors.Wrap(err, "unexpected query error")
		}
	}

	return nil
}
