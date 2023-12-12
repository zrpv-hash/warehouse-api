package getproducts

import (
	"warehousesvc/internal/application/product/getall"
)

type reqParams struct {
	Id   string `query:"id"`
	Page int    `query:"page"`
}

func (r *reqParams) toUsecasePayload() *getall.Payload {
	return &getall.Payload{
		WarehouseId: r.Id,
		Page:        r.Page,
	}
}

type responseBody struct {
	Products []Product `json:"products"`
}

type Product struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	Size        string `json:"size"`
	Quantity    int    `json:"quantity"`
}

func responseFromResult(r *getall.Result) *responseBody {
	res := make([]Product, len(*r))
	for k, v := range *r {
		res[k] = Product{
			ID:          v.ID,
			ProductName: v.ProductName,
			Size:        v.Size,
			Quantity:    v.Quantity,
		}
	}
	return &responseBody{res}
}
