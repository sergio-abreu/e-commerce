package inventory

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/inventory/application/inventory/proto"
	"github.com/sergio-vaz-abreu/inventory/modules/application"
	"google.golang.org/grpc"
)

func RegisterInventoryServer(server *grpc.Server, aInventory *application.Inventory) {
	proto.RegisterInventoryServer(server, &Inventory{aInventory: aInventory})
}

type Inventory struct {
	proto.InventoryServer
	aInventory *application.Inventory
}

func (i *Inventory) GetAllProducts(_ context.Context, _ *proto.Void) (*proto.Products, error) {
	products, err := i.aInventory.GetProducts()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get products")
	}
	var protoProducts []*proto.Product
	for _, aProduct := range products {
		protoProducts = append(protoProducts, &proto.Product{
			Id:           aProduct.Id,
			PriceInCents: int32(aProduct.PriceInCents),
			Title:        aProduct.Title,
			Description:  aProduct.Description,
		})
	}
	return &proto.Products{Products: protoProducts}, nil
}
