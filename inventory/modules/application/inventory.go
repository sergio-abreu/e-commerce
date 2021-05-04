package application

import (
	"github.com/sergio-vaz-abreu/inventory/modules/domain/product"
)

type Inventory struct {
	productsRepository product.Repository
}

func (i *Inventory) GetProducts() ([]product.Product, error) {
	return i.productsRepository.GetAll()
}
