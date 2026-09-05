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
	UpdateMaterial(materialStr *types.MaterialStr) (*types.Material, error)
	DeleteMaterial(id string) error
	ListProveedor(page int) (*types.ProveedorSlice, error)
	CreateProveedor(proveedorStr *types.ProveedorStr) (*types.Proveedor, error)
	ReadProveedor(id string) (*types.Proveedor, error)
	UpdateProveedor(proveedorStr *types.ProveedorStr) (*types.Proveedor, error)
	DeleteProveedor(id string) error
}

type service struct {
	store store.Store
}

func New(store store.Store) Service {
	return service{store: store}
}

// === === === MATERIAL === === ===

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

// === === === PROVEEDOR === === ===

func (s service) ListProveedor(page int) (*types.ProveedorSlice, error) {
	offset := (page - 1) * env.RecordsPerSlice
	records, count, err := s.store.ListProovedor(env.RecordsPerSlice, offset)
	if err != nil {
		return nil, err
	}
	totalPages := utils.CalculateTotalPages(count, env.RecordsPerSlice)
	hasNextPage := totalPages > page
	proveedorSlice := types.ProveedorSlice{
		Records:     records,
		HasNextPage: hasNextPage,
	}
	return &proveedorSlice, nil
}

func (s service) CreateProveedor(proveedorStr *types.ProveedorStr) (*types.Proveedor, error) {
	phone, err := strconv.Atoi(proveedorStr.Phone)
	if err != nil {
		return nil, err
	}
	proveedor := types.Proveedor{
		Name:  proveedorStr.Name,
		Phone: phone,
	}
	newProveedor, err := s.store.CreateProveedor(&proveedor)
	if err != nil {
		return nil, err
	}
	return newProveedor, nil
}

func (s service) ReadProveedor(id string) (*types.Proveedor, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	proveedor, err := s.store.ReadProveedor(intId)
	if err != nil {
		return nil, err
	}
	return proveedor, nil
}

func (s service) UpdateProveedor(proveedorStr *types.ProveedorStr) (*types.Proveedor, error) {
	intInd, err := strconv.Atoi(proveedorStr.Id)
	if err != nil {
		return nil, err
	}
	intPhone, err := strconv.Atoi(proveedorStr.Phone)
	if err != nil {
		return nil, err
	}
	proveedor := types.Proveedor{
		Id:    intInd,
		Name:  proveedorStr.Name,
		Phone: intPhone,
	}
	updProveedor, err := s.store.UpdateProveedor(&proveedor)
	if err != nil {
		return nil, err
	}
	return updProveedor, nil
}

func (s service) DeleteProveedor(id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if err := s.store.DeleteProveedor(intId); err != nil {
		return err
	}
	return nil
}

// SERVICES PRODUCT
