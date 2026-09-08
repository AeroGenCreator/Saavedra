package store

import (
	"Saavedra/service/Quote/types"
	"database/sql"
)

type Store interface {
	Many2One() (*types.Many2One, error)
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return store{db: db}
}

func (s store) Many2One() (*types.Many2One, error) {

	qUser := "SELECT id, name FROM users;"
	qCustomer := "SELECT id, name, customeremail FROM customer;"
	qProduct := `
	SELECT p.id, p.name, p.description, p.pmeasure, p.price, c.id, c.name, pr.id, pr.name
	FROM product AS p
	LEFT JOIN category AS c ON p.category_id = c.id
	LEFT JOIN proveedor AS pr ON p.proveedor_id = pr.id;`

	rowsUser, err := s.db.Query(qUser)
	if err != nil {
		return nil, err
	}
	defer rowsUser.Close()
	var userArray []*types.User
	for rowsUser.Next() {
		var record types.User
		if err = rowsUser.Scan(&record.Id, &record.Name); err != nil {
			return nil, err
		}
		userArray = append(userArray, &record)
	}
	if rowsUser.Err() != nil {
		return nil, err
	}

	rowsCustomer, err := s.db.Query(qCustomer)
	if err != nil {
		return nil, err
	}
	defer rowsCustomer.Close()
	var customerArray []*types.Customer
	for rowsCustomer.Next() {
		var record types.Customer
		if err = rowsCustomer.Scan(&record.Id, &record.Name, &record.CustomerEmail); err != nil {
			return nil, err
		}
		customerArray = append(customerArray, &record)
	}
	if rowsCustomer.Err() != nil {
		return nil, err
	}

	rowsProduct, err := s.db.Query(qProduct)
	if err != nil {
		return nil, err
	}
	defer rowsProduct.Close()
	var productArray []*types.ProductFetch
	for rowsProduct.Next() {
		var record types.ProductFetch
		if err = rowsProduct.Scan(
			&record.Id,
			&record.Name,
			&record.Description,
			&record.PMeasure,
			&record.Price,
			&record.CategoryId,
			&record.Category,
			&record.ProveedorId,
			&record.Proveedor,
		); err != nil {
			return nil, err
		}
		productArray = append(productArray, &record)
	}
	if rowsProduct.Err() != nil {
		return nil, err
	}

	return &types.Many2One{
		UserArray:     userArray,
		CustomerArray: customerArray,
		ProductArray:  productArray,
	}, nil
}
