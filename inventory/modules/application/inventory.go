package application

import (
	"context"
	"github.com/sergio-vaz-abreu/inventory/modules/domain/product"
	"go.elastic.co/apm"
)

func NewInventory(productsRepository product.Repository) *Inventory {
	return &Inventory{productsRepository: productsRepository}
}

type Inventory struct {
	productsRepository product.Repository
}

func (i *Inventory) GetProducts(ctx context.Context) ([]product.Product, error) {
	span, ctx := apm.StartSpan(ctx, "productsRepository.GetAll()", "runtime.product.repository")
	defer span.End()
	return i.productsRepository.GetAll()
}
