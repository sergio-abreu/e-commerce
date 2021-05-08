package application

import (
	"context"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"github.com/sergio-vaz-abreu/discount/application/discount_calculator/proto"
	"github.com/sergio-vaz-abreu/discount/infrastructure/postgres"
	"github.com/sergio-vaz-abreu/discount/infrastructure/time_seed"
	"github.com/sergio-vaz-abreu/discount/modules/domain/calculator"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestDiscountCalculator(t *testing.T) {
	g := NewGomegaWithT(t)
	address := ":50051"
	app, err := Load(GetConfig(address))
	g.Expect(err).Should(
		Not(HaveOccurred()))
	appErr := app.Run()
	g.Consistently(appErr).Should(
		Not(Receive()))
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	g.Expect(err).Should(
		Not(HaveOccurred()))
	t.Run("", func(t *testing.T) {
		time_seed.ChangeTimeSeed(func() time.Time {
			return calculator.BlackFriday
		})
		sut := proto.NewDiscountCalculatorClient(conn)

		discount, err := sut.CalculateUsersDiscountForProduct(context.Background(), &proto.RequestForDiscount{
			UserId:    "11054c65-89dd-46a6-86ab-c603195100a5",
			ProductId: "4151e87c-6a90-4f60-ae20-1643d5205fe3",
		})

		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(discount).Should(
			PointTo(MatchFields(IgnoreExtras, Fields{
				"Percentage":   BeEquivalentTo(0.1),
				"ValueInCents": BeEquivalentTo(13410),
			})))
		g.Consistently(appErr).Should(
			Not(Receive()))
	})
}

func GetConfig(address string) Config {
	return Config{
		PostgresConfig: postgres.Config{
			User:     "discount",
			Password: "P@ssword",
			Host:     "127.0.0.1",
			Port:     5433,
			Database: "discount",
		},
		RpcAddress: address,
	}
}
