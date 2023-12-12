package pgsql

import (
	"context"
	"database/sql"
	"warehousesvc/internal/domain/product"
	"warehousesvc/internal/domain/reserve"
	"warehousesvc/internal/infrastructure/tx/pgsqltx"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) reserve.Repository {
	return &repo{db}
}

func (r *repo) Reserve(ctx context.Context, quantity int, optionId string, warehouseId string) error {
	const updateReserveQuery = `
		WITH Reserved AS (
			UPDATE inventory
			SET quantity = quantity - $1
			WHERE option_id = $2
			AND warehouse_id = $3
			RETURNING *
		)
		INSERT INTO product_reservation (id, unique_code_id, warehouse_id, quantity_reserved, reservation_date)
		SELECT $4, $2, $3, $1, CURRENT_DATE
		FROM Reserved;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	_, err := q.ExecContext(ctx, updateReserveQuery, quantity, optionId, warehouseId, uuid.New().String())
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return errors.Wrap(product.ErrProductNotFound, "failed to update, not found")
		default:
			return errors.Wrap(err, "unexpected query error")
		}
	}

	return nil
}

func (r *repo) DeleteReserved(ctx context.Context, optionIds []string, warehouseId string) error {
	const deleteReserveQuery = `
		DELETE FROM product_reservation
		WHERE unique_code_id = ANY($1)
		AND warehouse_id = $2;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	_, err := q.ExecContext(ctx, deleteReserveQuery, pq.Array(optionIds), warehouseId)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return errors.Wrap(product.ErrProductNotFound, "failed to update, not found")
		default:
			return errors.Wrap(err, "unexpected query error")
		}
	}

	return nil
}
