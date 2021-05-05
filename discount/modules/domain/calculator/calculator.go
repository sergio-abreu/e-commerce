package calculator

import (
	"github.com/sergio-vaz-abreu/discount/infrastructure/time_seed"
	"github.com/sergio-vaz-abreu/discount/modules/domain/discount"
	"github.com/sergio-vaz-abreu/discount/modules/domain/product"
	"github.com/sergio-vaz-abreu/discount/modules/domain/user"
	"time"
)

var BlackFriday = time.Date(2021, time.November, 25, 0, 0, 0, 0, time.UTC)

const maxPercentage = 10 * discount.Percent

func IsBlackFriday() bool {
	return time_seed.IsSameDayOfTheYear(time_seed.Now(), BlackFriday)
}

func CalculateUserDiscounts(aUser user.User, aProduct product.Product) discount.Discount {
	aDiscount := discount.NewDiscount(aProduct.PriceInCents, maxPercentage)
	if aUser.IsBirthDay() {
		aDiscount.AddByPercentage(5 * discount.Percent)
	}
	if IsBlackFriday() {
		aDiscount.AddByPercentage(10 * discount.Percent)
	}
	return aDiscount
}
