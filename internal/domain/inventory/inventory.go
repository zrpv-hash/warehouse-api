package inventory

type Inventory struct {
	OptionId string
	Quantity int
}

func FormData(optionId string, quantity int) *Inventory {
	return &Inventory{
		optionId,
		quantity,
	}
}
