package application

import (
	"context"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"github.com/sergio-vaz-abreu/inventory/application/inventory/proto"
	"github.com/sergio-vaz-abreu/inventory/infrastructure/postgres"
	"google.golang.org/grpc"
	"testing"
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
		sut := proto.NewInventoryClient(conn)

		products, err := sut.GetAllProducts(context.Background(), &proto.Void{})

		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(products).Should(
			PointTo(MatchFields(IgnoreExtras, Fields{
				"Products": ContainElement(PointTo(MatchFields(IgnoreExtras, Fields{
					"Id":           Equal("ff5d8560-c82e-44f0-9947-930852ca5ecc"),
					"PriceInCents": BeEquivalentTo(4390),
					"Title":        Equal("Suporte para Notebook"),
					"Description":  Equal("Ideal para notebook ou tablet de qualquer tamanho"),
				}))),
			})))
		g.Consistently(appErr).Should(
			Not(Receive()))
	})
}

func GetConfig(address string) Config {
	return Config{
		PostgresConfig: postgres.Config{
			User:     "inventory",
			Password: "P@ssword",
			Host:     "127.0.0.1",
			Port:     5432,
			Database: "inventory",
		},
		Address: address,
	}
}
