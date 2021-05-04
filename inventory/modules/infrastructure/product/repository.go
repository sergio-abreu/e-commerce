package product

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/inventory/modules/domain/product"
)

func NewSqlRepository(sql *sql.DB) *SqlRepository {
	return &SqlRepository{sql: sql}
}

type SqlRepository struct {
	sql *sql.DB
}

func (s *SqlRepository) GetAll() ([]product.Product, error) {
	query := `SELECT id, price_in_cents, title, description from products`
	rows, err := s.sql.Query(query)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to query products table")
	}
	var products []product.Product
	for rows.Next() {
		var aProduct product.Product
		err := rows.Scan(&aProduct.Id, &aProduct.PriceInCents, &aProduct.Title, &aProduct.Description)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan products table")
		}
		products = append(products, aProduct)
	}
	return products, nil
}
