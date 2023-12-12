package getall

import (
	"warehousesvc/internal/core/usecase"
	"warehousesvc/internal/domain/product"
)

type Payload struct {
	WarehouseId string
	Page        int
}

type Result = []product.Product

type UseCase = usecase.UseCase[*Payload, Result]

type implementation struct {
	productRepo product.Repository
}

func New(
	productRepo product.Repository,
) UseCase {
	return &implementation{productRepo}
}
