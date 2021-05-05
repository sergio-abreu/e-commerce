package calculator

import (
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/discount/modules/domain/calculator"
	"github.com/sergio-vaz-abreu/discount/modules/domain/discount"
	"github.com/sergio-vaz-abreu/discount/modules/domain/product"
	"github.com/sergio-vaz-abreu/discount/modules/domain/user"
)

func NewCalculator(userRepository user.Repository, productRepository product.Repository) *Calculator {
	return &Calculator{userRepository: userRepository, productRepository: productRepository}
}

type Calculator struct {
	userRepository    user.Repository
	productRepository product.Repository
}

func (d *Calculator) CalculateDiscount(userId, productId string) (discount.Discount, error) {
	aUser, err := d.userRepository.GetById(userId)
	if err != nil {
		return discount.Discount{}, errors.Wrap(err, "failed to get user")
	}
	aProduct, err := d.productRepository.GetById(productId)
	if err != nil {
		return discount.Discount{}, errors.Wrap(err, "failed to get product")
	}
	aDiscount := calculator.CalculateUserDiscounts(aUser, aProduct)
	return aDiscount, nil
}
