package product

import "github.com/google/uuid"

func NewProduct(priceInCents int) Product {
	return Product{Id: uuid.New().String(), PriceInCents: priceInCents}
}

type Product struct {
	Id           string
	PriceInCents int
}
