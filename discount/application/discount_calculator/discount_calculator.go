package discount_calculator

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/discount/application/discount_calculator/proto"
	"github.com/sergio-vaz-abreu/discount/modules/application/calculator"
	"google.golang.org/grpc"
)

func RegisterDiscountCalculatorServer(server *grpc.Server, aCalculator *calculator.Calculator) {
	proto.RegisterDiscountCalculatorServer(server, &DiscountCalculator{aCalculator: aCalculator})
}

type DiscountCalculator struct {
	proto.DiscountCalculatorServer
	aCalculator *calculator.Calculator
}

func (s *DiscountCalculator) CalculateUsersDiscountForProduct(ctx context.Context, requestForDiscount *proto.RequestForDiscount) (*proto.Discount, error) {
	aDiscount, err := s.aCalculator.CalculateDiscount(requestForDiscount.UserId, requestForDiscount.ProductId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to calculate discount")
	}
	return &proto.Discount{Percentage: aDiscount.Percentage.Float64(), ValueInCents: int32(aDiscount.ValueInCents)}, nil
}
