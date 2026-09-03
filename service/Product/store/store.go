package store

import (
	"Saavedra/service/Product/types"
	"database/sql"
)

type Store interface {
	ListMaterial(limit, offset int) ([]*types.Material, int, error)
	CreateMaterial(material *types.Material) (*types.Material, error)
	ReadMaterial(id int) (*types.Material, error)
	UpdateMaterial(material *types.Material) (*types.Material, error)
	DeleteMaterial(id int) error
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
	qInsert := `
	INSERT INTO material (name) VALUES (?)
	ON CONFLICT(name) DO UPDATE SET name = excluded.name;
	`
	_, err := s.db.Exec(qInsert, material.Name)
	if err != nil {
		return nil, err
	}

	qSelect := `SELECT id, name FROM material WHERE name = ?;`
	var newMaterial types.Material
	err = s.db.QueryRow(qSelect, material.Name).Scan(&newMaterial.Id, &newMaterial.Name)
	if err != nil {
		return nil, err
	}

	return &newMaterial, nil
}

func (s store) ReadMaterial(id int) (*types.Material, error) {
	q := "SELECT id, name FROM material WHERE id = ?;"
	var record types.Material
	err := s.db.QueryRow(q, id).Scan(&record.Name, &record.Id)
	if err == sql.ErrNoRows {
		return nil, types.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return &record, nil
}

func (s store) UpdateMaterial(material *types.Material) (*types.Material, error) {
	q := "UPDATE material SET name = ? WHERE id = ?;"
	_, err := s.db.Exec(q, material.Name, material.Id)
	if err != nil {
		return nil, err
	}
	return material, nil
}

func (s store) DeleteMaterial(id int) error {
	q := "DELETE FROM material WHERE id = ?;"
	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}
	return nil
}

// SECTION PROVEEDOR
// SECTION PRODUCT
