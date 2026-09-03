package store

import (
	"Saavedra/service/Product/types"
	"database/sql"
)

type Store interface {
	CreateMaterial(material *types.Material) (*types.Material, error)
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return store{db: db}
}

// SECTION MATERIALS
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
