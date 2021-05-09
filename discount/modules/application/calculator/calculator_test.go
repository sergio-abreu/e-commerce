package calculator

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"github.com/sergio-vaz-abreu/discount/infrastructure/postgres"
	"github.com/sergio-vaz-abreu/discount/infrastructure/time_seed"
	"github.com/sergio-vaz-abreu/discount/modules/domain/calculator"
	"github.com/sergio-vaz-abreu/discount/modules/infrastructure/product"
	"github.com/sergio-vaz-abreu/discount/modules/infrastructure/user"
	"testing"
	"time"
)

func TestCalculator(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("Get max discount for black friday and user's birthday", func(t *testing.T) {
		time_seed.ChangeTimeSeed(func() time.Time {
			return calculator.BlackFriday
		})
		sut, err := GetCalculator()
		g.Expect(err).Should(
			Not(HaveOccurred()))

		discount, err := sut.CalculateDiscount("11054c65-89dd-46a6-86ab-c603195100a5", "4151e87c-6a90-4f60-ae20-1643d5205fe3")

		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(discount).Should(
			MatchFields(IgnoreExtras, Fields{
				"Percentage":   BeEquivalentTo(0.1),
				"ValueInCents": BeEquivalentTo(13410),
			}))
	})
}

func GetCalculator() (*Calculator, error) {
	db, err := postgres.NewDatabase(postgres.Config{
		User:     "discount",
		Password: "P@ssword",
		Host:     "127.0.0.1",
		Port:     15432,
		Database: "discount",
	})
	if err != nil {
		return nil, err
	}
	userRepository := user.NewSqlRepository(db)
	productRepository := product.NewSqlRepository(db)
	return NewCalculator(userRepository, productRepository), nil
}
