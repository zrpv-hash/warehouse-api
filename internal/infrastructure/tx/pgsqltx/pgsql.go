package pgsqltx

import (
	"context"
	"warehousesvc/internal/core/tx"
	"warehousesvc/internal/infrastructure/tx/isolation"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type pgsqlTx struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) tx.TransactionManager {
	return &pgsqlTx{db}
}

func (tx *pgsqlTx) Do(ctx context.Context, level isolation.IsolationLevel, fn func(context.Context) error) (err error) {
	ctx, commit, err := tx.wrapContextAndGetCommitFn(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to wrap context and start transaction")
	}

	defer commit(&err)

	_, err = tx.db.ExecContext(ctx, "SET TRANSACTION ISOLATION LEVEL "+level.String())
	if err != nil {
		return errors.Wrap(err, "failed to set transaction isolation level")
	}

	if err := fn(ctx); err != nil {
		return errors.Wrap(err, "failed to execute transactional fn")
	}

	return nil
}
