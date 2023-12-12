package tx

import (
	"context"
	"errors"
	"warehousesvc/internal/infrastructure/tx/isolation"
)

var (
	ErrUnexpetedTxFailure = errors.New("transaction unexpected failure")
)

type TransactionManager interface {
	Do(context.Context, isolation.IsolationLevel, func(context.Context) error) error
}
