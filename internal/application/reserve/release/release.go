package release

import (
	"warehousesvc/internal/core/tx"
	"warehousesvc/internal/core/usecase"
	"warehousesvc/internal/domain/inventory"
	"warehousesvc/internal/domain/reserve"
)

type Payload struct {
	WarehouseId string
	UniqueIds   []string
}

func (p *Payload) GetArray() []string {
	return p.UniqueIds
}

type UseCase = usecase.Interactor[*Payload]

type implementation struct {
	txManager     tx.TransactionManager
	inventoryRepo inventory.Repository
	reserveRepo   reserve.Repository
}

func New(
	txManager tx.TransactionManager,
	inventoryRepo inventory.Repository,
	reserveRepo reserve.Repository,
) UseCase {
	return &implementation{txManager, inventoryRepo, reserveRepo}
}
