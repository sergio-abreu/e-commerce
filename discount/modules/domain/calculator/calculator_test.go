package calculator

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"github.com/sergio-vaz-abreu/discount/infrastructure/time_seed"
	"github.com/sergio-vaz-abreu/discount/modules/domain/product"
	"github.com/sergio-vaz-abreu/discount/modules/domain/user"
	"testing"
	"time"
)

func TestCalculateUserDiscounts(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("Calculating 5% of discount when it is user birthday", func(t *testing.T) {
		birthDay := time.Date(2000, time.May, 04, 0, 0, 0, 0, time.UTC)
		time_seed.ChangeTimeSeed(func() time.Time {
			return birthDay.AddDate(21, 0, 0)
		})
		aUser := user.NewUser(birthDay)
		aProduct := product.NewProduct(100)

		discount := CalculateUserDiscounts(aUser, aProduct)

		g.Expect(discount).Should(
			MatchFields(IgnoreExtras, Fields{
				"Percentage":   BeEquivalentTo(0.05),
				"ValueInCents": BeEquivalentTo(95),
			}))
	})
	t.Run("Calculating 10% of discount when it is black friday", func(t *testing.T) {
		birthDay := time.Date(2000, time.May, 04, 0, 0, 0, 0, time.UTC)
		time_seed.ChangeTimeSeed(func() time.Time {
			return BlackFriday
		})
		aUser := user.NewUser(birthDay)
		aProduct := product.NewProduct(100)

		discount := CalculateUserDiscounts(aUser, aProduct)

		g.Expect(discount).Should(
			MatchFields(IgnoreExtras, Fields{
				"Percentage":   BeEquivalentTo(0.1),
				"ValueInCents": BeEquivalentTo(90),
			}))
	})
	t.Run("Calculating 10% of discount when it is black friday and user's birthday", func(t *testing.T) {
		birthDay := BlackFriday
		time_seed.ChangeTimeSeed(func() time.Time {
			return birthDay
		})
		aUser := user.NewUser(birthDay)
		aProduct := product.NewProduct(100)

		discount := CalculateUserDiscounts(aUser, aProduct)

		g.Expect(discount).Should(
			MatchFields(IgnoreExtras, Fields{
				"Percentage":   BeEquivalentTo(0.1),
				"ValueInCents": BeEquivalentTo(90),
			}))
	})
}
