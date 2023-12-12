package deleterelease

import (
	"warehousesvc/internal/application/reserve/release"
)

type requestBody struct {
	WarehouseId string   `json:"warehouseId"`
	UniqueIds   []string `json:"uniqueIds"`
}

func (r *requestBody) toUsecasePayload() *release.Payload {
	return &release.Payload{
		WarehouseId: r.WarehouseId,
		UniqueIds:   r.UniqueIds,
	}
}
