package calculator

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/discount/modules/domain/calculator"
	"github.com/sergio-vaz-abreu/discount/modules/domain/discount"
	"github.com/sergio-vaz-abreu/discount/modules/domain/product"
	"github.com/sergio-vaz-abreu/discount/modules/domain/user"
	"go.elastic.co/apm"
)

func NewCalculator(userRepository user.Repository, productRepository product.Repository) *Calculator {
	return &Calculator{userRepository: userRepository, productRepository: productRepository}
}

type Calculator struct {
	userRepository    user.Repository
	productRepository product.Repository
}

func (d *Calculator) CalculateDiscount(ctx context.Context, userId, productId string) (discount.Discount, error) {
	span, ctx := apm.StartSpan(ctx, "CalculateDiscount()", "runtime.calculator")
	defer span.End()
	userRepoSpan, _ := apm.StartSpan(ctx, "userRepository.GetById()", "runtime.user.repository")
	aUser, err := d.userRepository.GetById(userId)
	userRepoSpan.End()
	if err != nil {
		return discount.Discount{}, errors.Wrap(err, "failed to get user")
	}
	productRepoSpan, _ := apm.StartSpan(ctx, "productRepository.GetById()", "runtime.product.repository")
	aProduct, err := d.productRepository.GetById(productId)
	productRepoSpan.End()
	if err != nil {
		return discount.Discount{}, errors.Wrap(err, "failed to get product")
	}
	calculatorSpan, _ := apm.StartSpan(ctx, "calculator.CalculateUserDiscounts()", "runtime.calculator.calculate")
	aDiscount := calculator.CalculateUserDiscounts(aUser, aProduct)
	calculatorSpan.End()
	return aDiscount, nil
}
