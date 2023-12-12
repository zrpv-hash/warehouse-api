package postreserve

import (
	"warehousesvc/internal/application/reserve/reserve"
)

type requestBody struct {
	WarehouseId  string         `json:"warehouseId"`
	ReserveItems []reserveItems `json:"reserveItems" validate:"dive"`
}

type reserveItems struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity" validate:"required,min=1"`
}

func (r *requestBody) toUsecasePayload() *reserve.Payload {
	p := reserve.Payload{UniqueCodeMap: make(map[string]int, len(r.ReserveItems)), WarehouseId: r.WarehouseId}

	for _, v := range r.ReserveItems {
		p.UniqueCodeMap[v.ID] = v.Quantity
	}
	return &p
}
