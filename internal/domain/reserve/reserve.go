package reserve

type Inventory struct {
	OptionId, WarehouseId string
	Quantity              int
}

func FormData(optionId string, warehouseId string, quantity int) *Inventory {
	return &Inventory{
		optionId,
		warehouseId,
		quantity,
	}
}
