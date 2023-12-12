package reserve

import "context"

type Repository interface {
	Reserve(context.Context, int, string, string) error
	DeleteReserved(context.Context, []string, string) error
}
