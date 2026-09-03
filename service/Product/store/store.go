package store

import (
	"Saavedra/service/Product/types"
	"database/sql"
)

type Store interface {
	CreateMaterial(material *types.Material) (*types.Material, error)
	ListMaterial(limit, offset int) ([]*types.Material, int, error)
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return store{db: db}
}

// SECTION MATERIALS
func (s store) ListMaterial(limit, offset int) ([]*types.Material, int, error) {
	q1 := "SELECT COUNT(id) AS count_id FROM material;"
	q2 := "SELECT id, name FROM material LIMIT ? OFFSET ?;"

	var count int
	if err := s.db.QueryRow(q1).Scan(&count); err != nil {
		return nil, 0, err
	}
	rows, err := s.db.Query(q2, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var records []*types.Material
	for rows.Next() {
		var record types.Material
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

func (s store) CreateMaterial(material *types.Material) (*types.Material, error) {
	q := `
	INSERT INTO material (name) VALUES (?)
	ON CONFLICT(name) DO UPDATE SET name = excluded.name
	RETURNING id, name;
	`
	var newMaterial types.Material
	err := s.db.QueryRow(q, material.Name).Scan(&newMaterial.Id, &newMaterial.Name)
	if err != nil {
		return nil, err
	}

	return &newMaterial, nil
}

// SECTION PROVEEDOR
// SECTION PRODUCT
