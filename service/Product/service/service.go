package service

import (
	"Saavedra/env"
	"Saavedra/service/Product/store"
	"Saavedra/service/Product/types"
	"Saavedra/utils"
	"strconv"
)

type Service interface {
	ListMaterial(page int) (*types.MaterialSlice, error)
	CreateMaterial(material *types.Material) (*types.Material, error)
	ReadMaterial(id string) (*types.Material, error)
	UpdateMaterial(material *types.MaterialStr) (*types.Material, error)
	DeleteMaterial(id string) error
}

type service struct {
	store store.Store
}

func New(store store.Store) Service {
	return service{store: store}
}

// SERVICES MATERIAL
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

func (s service) CreateMaterial(material *types.Material) (*types.Material, error) {
	newMaterial, err := s.store.CreateMaterial(material)
	if err != nil {
		return nil, err
	}
	return newMaterial, nil
}

func (s service) ReadMaterial(id string) (*types.Material, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	material, err := s.store.ReadMaterial(intId)
	if err != nil {
		return nil, err
	}
	return material, nil
}

func (s service) UpdateMaterial(materialStr *types.MaterialStr) (*types.Material, error) {
	intInd, err := strconv.Atoi(materialStr.Id)
	if err != nil {
		return nil, err
	}
	material := types.Material{
		Id:   intInd,
		Name: materialStr.Name,
	}
	updMaterial, err := s.store.UpdateMaterial(&material)
	if err != nil {
		return nil, err
	}
	return updMaterial, nil
}

func (s service) DeleteMaterial(id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if err := s.store.DeleteMaterial(intId); err != nil {
		return err
	}
	return nil
}

// SERVICES PROVEEDOR
// SERVICES PRODUCT
