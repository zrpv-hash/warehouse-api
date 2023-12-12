package reserve

import (
	"warehousesvc/internal/core/tx"
	"warehousesvc/internal/core/usecase"
	"warehousesvc/internal/domain/inventory"
	"warehousesvc/internal/domain/reserve"
)

type Payload struct {
	WarehouseId   string
	UniqueCodeMap map[string]int
}

func (p *Payload) GetArray() []string {
	res := make([]string, len(p.UniqueCodeMap))
	for v := range p.UniqueCodeMap {
		res = append(res, v)
	}
	return res
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
