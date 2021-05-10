package inventory

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/inventory/application/inventory/proto"
	"github.com/sergio-vaz-abreu/inventory/modules/application"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm"
	"google.golang.org/grpc"
)

func RegisterInventoryServer(server *grpc.Server, aInventory *application.Inventory) {
	proto.RegisterInventoryServer(server, &Inventory{aInventory: aInventory})
}

type Inventory struct {
	proto.InventoryServer
	aInventory *application.Inventory
}

func (i *Inventory) GetAllProducts(ctx context.Context, _ *proto.Void) (*proto.Products, error) {
	logger := logrus.WithContext(ctx)
	span, ctx := apm.StartSpan(ctx, "GetAllProducts()", "runtime.inventory")
	defer span.End()
	products, err := i.aInventory.GetProducts(ctx)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		logger.WithError(err).Error("failed to get products")
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
	logger.WithError(err).Error("successfully got products")
	return &proto.Products{Products: protoProducts}, nil
}
