package product

type Product struct {
	ID string

	ProductName, Size string

	Quantity int
}

type Products struct {
	Products []Product
}

func FormData(id, productName, size string, quantity int) *Product {
	return &Product{
		id,
		productName,
		size,
		quantity,
	}
}
