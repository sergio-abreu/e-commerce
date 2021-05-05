package product

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/discount/modules/domain/product"
)

var ErrProductNotFound = errors.New("product not found")

type Repository struct {
	db *sql.DB
}

func (r *Repository) GetById(id string) (product.Product, error) {
	query := "SELECT id, price_in_cents FROM products WHERE id = $1"
	row := r.db.QueryRow(query, id)
	var aProduct product.Product
	err := row.Scan(&aProduct.Id, &aProduct.PriceInCents)
	if err == sql.ErrNoRows {
		return aProduct, ErrProductNotFound
	}
	if err != nil {
		return aProduct, errors.Wrap(err, "failed to query product database")
	}
	return aProduct, nil
}

func (r *Repository) GetAll() ([]product.Product, error) {
	query := "SELECT id, price_in_cents FROM products"
	rows, err := r.db.Query(query)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to query product database")
	}
	var products []product.Product
	for rows.Next() {
		var aProduct product.Product
		err := rows.Scan(&aProduct.Id, &aProduct.PriceInCents)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan product database")
		}
		products = append(products, aProduct)
	}
	return products, nil
}
