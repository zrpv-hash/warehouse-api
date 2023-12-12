package pgsqltx

import (
	"context"
	"warehousesvc/internal/infrastructure/crosscutting/pgclient"

	"github.com/jmoiron/sqlx"
)

type querierCtxKey struct{}

func QuerierFromCtx(ctx context.Context, db *sqlx.DB) pgclient.Querier {
	if tr := querierFromCtx(ctx); tr != nil {
		return tr
	}

	if db != nil {
		return pgclient.Querier(db)
	}

	return nil
}

func IsQuerierInContext(ctx context.Context) bool {
	return querierFromCtx(ctx) != nil
}

func (t *pgsqlTx) wrapContextAndGetCommitFn(ctx context.Context) (context.Context, commitFn, error) {
	if q := QuerierFromCtx(ctx, nil); q == nil {
		tx, err := t.db.Beginx()
		if err != nil {
			return nil, nil, err
		}

		return ctxWithQuerier(ctx, tx), sqlxCommit(tx), nil
	}

	return ctx, noopCommit(), nil
}

func querierFromCtx(ctx context.Context) pgclient.Querier {
	if tr, ok := ctx.Value(querierCtxKey{}).(pgclient.Querier); ok {
		return tr
	}

	return nil
}

func ctxWithQuerier(ctx context.Context, q pgclient.Querier) context.Context {
	return context.WithValue(ctx, querierCtxKey{}, q)
}
