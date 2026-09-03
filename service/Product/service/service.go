package service

import (
	"Saavedra/service/Product/store"
	"Saavedra/service/Product/types"
)

type Service interface {
	CreateMaterial(material *types.Material) (*types.Material, error)
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

// SERVICES PROVEEDOR
// SERVICES PRODUCT
