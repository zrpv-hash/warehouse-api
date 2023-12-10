package pgclient

import (
	"context"
	"fmt"
	"time"
	"warehousesvc/internal/core/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Querier interface {
	sqlx.QueryerContext
	sqlx.ExecerContext
	sqlx.ExtContext

	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

const connTimeout = time.Second * 3

func New(config *config.Config) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "postgres", buildConnectionString(config))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to postgres")
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to ping postgres")
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	return db, nil
}

func buildConnectionString(config *config.Config) string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)
}
