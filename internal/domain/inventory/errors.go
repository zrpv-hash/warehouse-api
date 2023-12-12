package inventory

import "warehousesvc/internal/core/domainerr"

var (
	ErrProductNotFound = domainerr.New("product not found")
)
