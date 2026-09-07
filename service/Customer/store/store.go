package store

import (
	"Saavedra/service/Customer/types"
	"database/sql"
)

type Store interface {
	ListCustomer(limit, offset int) ([]*types.Customer, int, error)
	CreateCustomer(customer types.Customer) error
	ReadCustomer(id int) (*types.Customer, error)
	UpdateCustomer(customer *types.Customer) error
	DeleteCustomer(id int) error
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return store{db: db}
}

func (s store) ListCustomer(limit, offset int) ([]*types.Customer, int, error) {
	q := `
	SELECT id, name, fullname, address, technicianphone, buyerphone, customeremail
	FROM customer
	LIMIT ? OFFSET ?;`
	q2 := "SELECT COUNT(id) AS count_id FROM customer;"
	rows, err := s.db.Query(q, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var customerSlice []*types.Customer
	for rows.Next() {
		var record types.Customer
		err = rows.Scan(
			&record.Id,
			&record.Name,
			&record.FullName,
			&record.Address,
			&record.TechnicianPhone,
			&record.BuyerPhone,
			&record.CustomerEmail)
		if err != nil {
			return nil, 0, err
		}
		customerSlice = append(customerSlice, &record)
	}
	if rows.Err() != nil {
		return nil, 0, err
	}
	var count int
	if err = s.db.QueryRow(q2).Scan(&count); err != nil {
		return nil, 0, err
	}
	return customerSlice, count, nil
}

func (s store) CreateCustomer(customer types.Customer) error {
	q := `
	INSERT INTO customer
	(name, fullname, address, technicianphone, buyerphone, customeremail)
	VALUES (?, ?, ?, ?, ?, ?)
	ON CONFLICT(name, fullname)
	DO UPDATE SET
	address = excluded.address
	technicianphone = excluded.technicianphone
	buyerphone = excluded.buyerphone
	customeremail = excluded.customeremail;`
	_, err := s.db.Exec(
		q,
		customer.Name,
		customer.FullName,
		customer.Address,
		customer.TechnicianPhone,
		customer.BuyerPhone,
		customer.CustomerEmail)
	if err != nil {
		return err
	}
	return nil
}

func (s store) ReadCustomer(id int) (*types.Customer, error) {
	q := `
	SELECT id, name, fullname, address, technicianphone, buyerphone, customeremail
	FROM customer
	WHERE id = ?;`
	var record types.Customer
	if err := s.db.QueryRow(q, id).Scan(
		&record.Id,
		&record.Name,
		&record.FullName,
		&record.Address,
		&record.TechnicianPhone,
		&record.BuyerPhone,
		&record.CustomerEmail); err != nil {
		return nil, err
	}
	return &record, nil
}

func (s store) UpdateCustomer(customer *types.Customer) error {
	q := `
	UPDATE customer SET
	name = ?,
	fullname = ?,
	address = ?,
	technicianphone = ?,
	buyerphone = ?,
	customeremail = ?
	WHERE id = ?;`
	_, err := s.db.Exec(
		q,
		customer.Name,
		customer.FullName,
		customer.Address,
		customer.TechnicianPhone,
		customer.BuyerPhone,
		customer.CustomerEmail)
	if err != nil {
		return err
	}
	return nil
}

func (s store) DeleteCustomer(id int) error {
	q := `DELETE FROM customer WHERE id = ?;`
	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}
	return nil
}
