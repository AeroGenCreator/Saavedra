package service

import (
	"Saavedra/env"
	"Saavedra/service/Product/store"
	"Saavedra/service/Product/types"
	"Saavedra/utils"
)

type Service interface {
	ListMaterial(page int) (*types.MaterialSlice, error)
	CreateMaterial(material *types.Material) (*types.Material, error)
	ReadMaterial(id string) (*types.Material, error)
	UpdateMaterial(material *types.Material) (*types.Material, error)
	DeleteMaterial(id string) error
}

type service struct {
	store store.Store
}

func New(store store.Store) Service {
	return service{store: store}
}

// SERVICES MATERIAL
func (s service) CreateMaterial(material *types.Material) (*types.Material, error) {
	newMaterial, err := s.store.CreateMaterial(material)
	if err != nil {
		return nil, err
	}
	return newMaterial, nil
}

func (s service) ListMaterial(page int) (*types.MaterialSlice, error) {
	offset := (page - 1) * env.RecordsPerSlice
	records, count, err := s.store.ListMaterial(env.RecordsPerSlice, offset)
	if err != nil {
		return nil, err
	}
	totalPages := utils.CalculateTotalPages(count, env.RecordsPerSlice)
	hasNextPage := totalPages > page
	materialSlice := types.MaterialSlice{
		Records:     records,
		HasNextPage: hasNextPage,
	}
	return &materialSlice, nil
}

func (s service) ReadMaterial(id string) (*types.Material, error) {
	// transformar id a int
	return nil, nil
}

func (s service) UpdateMaterial(material *types.Material) (*types.Material, error) {
	return nil, nil
}

func (s service) DeleteMaterial(id string) error {
	return nil
}

// SERVICES PROVEEDOR
// SERVICES PRODUCT
