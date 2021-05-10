package discount_calculator

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/discount/application/discount_calculator/proto"
	"github.com/sergio-vaz-abreu/discount/modules/application/calculator"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm"
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
	logger := logrus.WithContext(ctx)
	span, ctx := apm.StartSpan(ctx, "CalculateUsersDiscountForProduct", "runtime.discount")
	defer span.End()
	aDiscount, err := s.aCalculator.CalculateDiscount(ctx, requestForDiscount.UserId, requestForDiscount.ProductId)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		logger.WithContext(ctx).WithError(err).Error("failed to calculate discount")
		return nil, errors.Wrap(err, "failed to calculate discount")
	}
	logger.WithContext(ctx).Info("successfully calculated discount")
	return &proto.Discount{Percentage: aDiscount.Percentage.Float64(), ValueInCents: int32(aDiscount.ValueInCents)}, nil
}
