package store

import (
	"Saavedra/service/Product/types"
	"database/sql"
)

type Store interface {
	LoadAllCategory() ([]*types.Category, error)
	LoadAllProveedor() ([]*types.Proveedor, error)
	ListCategory(limit, offset int) ([]*types.Category, int, error)
	CreateCategory(category *types.Category) (*types.Category, error)
	ReadCategory(id int) (*types.Category, error)
	UpdateCategory(category *types.Category) (*types.Category, error)
	DeleteCategory(id int) error
	ListProovedor(limit, offset int) ([]*types.Proveedor, int, error)
	CreateProveedor(proveedor *types.Proveedor) (*types.Proveedor, error)
	ReadProveedor(id int) (*types.Proveedor, error)
	UpdateProveedor(proveedor *types.Proveedor) (*types.Proveedor, error)
	DeleteProveedor(id int) error
	ListProduct(limit, offset int) ([]*types.ProductFetch, int, error)
	CreateProduct(product *types.Product) error
	ReadProduct(id int) (*types.ProductFetch, error)
	UpdateProduct(product *types.Product) error
	DeleteProduct(id int) error
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return store{db: db}
}

func (s store) LoadAllCategory() ([]*types.Category, error) {
	q := "SELECT id, name FROM category;"
	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categoryArray []*types.Category
	for rows.Next() {
		var record types.Category
		if err := rows.Scan(&record.Id, &record.Name); err != nil {
			return nil, err
		}
		categoryArray = append(categoryArray, &record)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return categoryArray, nil
}

func (s store) LoadAllProveedor() ([]*types.Proveedor, error) {
	q := "SELECT id, name FROM proveedor;"
	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var proveedorArray []*types.Proveedor
	for rows.Next() {
		var record types.Proveedor
		if err := rows.Scan(&record.Id, &record.Name); err != nil {
			return nil, err
		}
		proveedorArray = append(proveedorArray, &record)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return proveedorArray, nil
}

// === === === CATEGORY === === ===

func (s store) ListCategory(limit, offset int) ([]*types.Category, int, error) {
	q1 := "SELECT COUNT(id) AS count_id FROM category;"
	q2 := "SELECT id, name FROM category LIMIT ? OFFSET ?;"

	var count int
	if err := s.db.QueryRow(q1).Scan(&count); err != nil {
		return nil, 0, err
	}
	rows, err := s.db.Query(q2, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var records []*types.Category
	for rows.Next() {
		var record types.Category
		err = rows.Scan(&record.Id, &record.Name)
		if err != nil {
			return nil, 0, err
		}
		records = append(records, &record)
	}
	if rows.Err() != nil {
		return nil, 0, rows.Err()
	}
	return records, count, nil
}

func (s store) CreateCategory(category *types.Category) (*types.Category, error) {
	qInsert := `
	INSERT INTO category (name) VALUES (?)
	ON CONFLICT(name) DO UPDATE SET name = excluded.name;
	`
	_, err := s.db.Exec(qInsert, category.Name)
	if err != nil {
		return nil, err
	}

	qSelect := `SELECT id, name FROM category WHERE name = ?;`
	var newCategory types.Category
	err = s.db.QueryRow(qSelect, category.Name).Scan(&newCategory.Id, &newCategory.Name)
	if err != nil {
		return nil, err
	}

	return &newCategory, nil
}

func (s store) ReadCategory(id int) (*types.Category, error) {
	q := "SELECT id, name FROM category WHERE id = ?;"
	var record types.Category
	err := s.db.QueryRow(q, id).Scan(&record.Id, &record.Name)
	if err == sql.ErrNoRows {
		return nil, types.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return &record, nil
}

func (s store) UpdateCategory(category *types.Category) (*types.Category, error) {
	q := "UPDATE category SET name = ? WHERE id = ?;"
	_, err := s.db.Exec(q, category.Name, category.Id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s store) DeleteCategory(id int) error {
	q := "DELETE FROM category WHERE id = ?;"
	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}
	return nil
}

// === === === PROVEEDOR === === ===

func (s store) ListProovedor(limit, offset int) ([]*types.Proveedor, int, error) {
	q1 := "SELECT COUNT(id) AS count_id FROM proveedor;"
	q2 := "SELECT id, name, phone FROM proveedor LIMIT ? OFFSET ?;"

	var count int
	if err := s.db.QueryRow(q1).Scan(&count); err != nil {
		return nil, 0, err
	}
	rows, err := s.db.Query(q2, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var records []*types.Proveedor
	for rows.Next() {
		var record types.Proveedor
		err = rows.Scan(&record.Id, &record.Name, &record.Phone)
		if err != nil {
			return nil, 0, err
		}
		records = append(records, &record)
	}
	if rows.Err() != nil {
		return nil, 0, rows.Err()
	}
	return records, count, nil
}

func (s store) CreateProveedor(proveedor *types.Proveedor) (*types.Proveedor, error) {
	qInsert := `
	INSERT INTO proveedor (name, phone) VALUES (?, ?)
	ON CONFLICT(name) DO UPDATE SET phone = excluded.phone;
	`
	_, err := s.db.Exec(qInsert, proveedor.Name, proveedor.Phone)
	if err != nil {
		return nil, err
	}

	qSelect := `SELECT id, name, phone FROM proveedor WHERE name = ?;`
	var newProveedor types.Proveedor
	err = s.db.QueryRow(qSelect, proveedor.Name).Scan(&newProveedor.Id, &newProveedor.Name, &newProveedor.Phone)
	if err != nil {
		return nil, err
	}

	return &newProveedor, nil
}

func (s store) ReadProveedor(id int) (*types.Proveedor, error) {
	q := "SELECT id, name, phone FROM proveedor WHERE id = ?;"
	var record types.Proveedor
	err := s.db.QueryRow(q, id).Scan(&record.Id, &record.Name, &record.Phone)
	if err == sql.ErrNoRows {
		return nil, types.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return &record, nil
}

func (s store) UpdateProveedor(proveedor *types.Proveedor) (*types.Proveedor, error) {
	q := "UPDATE proveedor SET name = ?, phone = ? WHERE id = ?;"
	_, err := s.db.Exec(q, proveedor.Name, proveedor.Phone, proveedor.Id)
	if err != nil {
		return nil, err
	}
	return proveedor, nil
}

func (s store) DeleteProveedor(id int) error {
	q := "DELETE FROM proveedor WHERE id = ?;"
	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}
	return nil
}

// === === === PRODUCT === === ===

func (s store) ListProduct(limit, offset int) ([]*types.ProductFetch, int, error) {
	q1 := "SELECT COUNT(id) AS count_id FROM product;"
	q2 := `SELECT p.id, p.name, p.description, p.pmeasure, p.price, m.id, m.name, pr.id, pr.name
	FROM product AS p
	LEFT JOIN category AS m ON p.category_id = m.id
	LEFT JOIN proveedor AS pr ON p.proveedor_id = pr.id
	LIMIT ? OFFSET ?;`
	var count int
	if err := s.db.QueryRow(q1).Scan(&count); err != nil {
		return nil, 0, err
	}
	rows, err := s.db.Query(q2, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var records []*types.ProductFetch
	for rows.Next() {
		var record types.ProductFetch
		if err = rows.Scan(
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
			return nil, 0, err
		}
		records = append(records, &record)
	}
	if rows.Err() != nil {
		return nil, 0, rows.Err()
	}
	return records, count, nil
}

func (s store) CreateProduct(product *types.Product) error {
	q1 := `INSERT INTO product (
	name, description, pmeasure, price, category_id, proveedor_id)
	VALUES (?, ?, ?, ?, ?, ?)
	ON CONFLICT(name, pmeasure, proveedor_id)
	DO UPDATE SET
	description = excluded.description,
	price = excluded.price,
	category_id = excluded.category_id;`
	_, err := s.db.Exec(
		q1,
		product.Name,
		product.Description,
		product.PMeasure,
		product.Price,
		product.CategoryId,
		product.ProveedorId)
	if err != nil {
		return err
	}
	return nil
}

func (s store) ReadProduct(id int) (*types.ProductFetch, error) {
	q := `SELECT p.id, p.name, p.description, p.pmeasure, p.price, m.id, m.name, pr.id, pr.name
	FROM product AS p
	LEFT JOIN category AS m ON p.category_id = m.id
	LEFT JOIN proveedor AS pr ON p.proveedor_id = pr.id
	WHERE p.id = ?;`
	var record types.ProductFetch
	err := s.db.QueryRow(q, id).Scan(
		&record.Id,
		&record.Name,
		&record.Description,
		&record.PMeasure,
		&record.Price,
		&record.CategoryId,
		&record.Category,
		&record.ProveedorId,
		&record.Proveedor,
	)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (s store) UpdateProduct(product *types.Product) error {
	q := `
	UPDATE product
	SET name = ?, description = ?, pmeasure = ?, price = ?, category_id = ?, proveedor_id = ?
	WHERE id = ?;`
	_, err := s.db.Exec(
		q,
		product.Name,
		product.Description,
		product.PMeasure,
		product.Price,
		product.CategoryId,
		product.ProveedorId,
		product.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s store) DeleteProduct(id int) error {
	q := "DELETE FROM product WHERE id = ?;"
	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}
	return nil
}
