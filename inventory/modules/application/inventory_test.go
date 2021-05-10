package application

import (
	"context"
	"database/sql"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"github.com/sergio-vaz-abreu/inventory/infrastructure/postgres"
	"github.com/sergio-vaz-abreu/inventory/modules/infrastructure/product"
	"testing"
)

func TestInventory(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("Get All products", func(t *testing.T) {
		sut, err := GetInventory()
		g.Expect(err).Should(
			Not(HaveOccurred()))

		products, err := sut.GetProducts(context.Background())

		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(products).Should(
			ContainElement(MatchAllFields(Fields{
				"Id":           Equal("ff5d8560-c82e-44f0-9947-930852ca5ecc"),
				"PriceInCents": BeEquivalentTo(4390),
				"Title":        Equal("Suporte para Notebook"),
				"Description":  Equal("Ideal para notebook ou tablet de qualquer tamanho"),
			})))
	})
}

func GetInventory() (*Inventory, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}
	repository := product.NewSqlRepository(db)
	inventory := NewInventory(repository)
	return inventory, nil
}

func GetDatabase() (*sql.DB, error) {
	db, err := postgres.NewDatabase(postgres.Config{
		User:     "inventory",
		Password: "P@ssword",
		Host:     "127.0.0.1",
		Port:     25432,
		Database: "inventory",
	})
	return db, err
}
