package service

import (
	"Saavedra/env"
	"Saavedra/service/Customer/store"
	"Saavedra/service/Customer/types"
	"Saavedra/utils"
	"strconv"
)

type Service interface {
	ListCustomer(page string) (*types.CustomerSlice, error)
	CreateCustomer(customer *types.Customer) error
	ReadCustomer(id string) (*types.Customer, error)
	UpdateCustomer(customer *types.Customer) error
	DeleteCustomer(id string) error
}

type service struct {
	store store.Store
}

func New(store store.Store) Service {
	return service{store: store}
}

func (s service) ListCustomer(page string) (*types.CustomerSlice, error) {
	intPage, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	offset := (intPage - 1) * env.RecordsPerSlice
	records, count, err := s.store.ListCustomer(env.RecordsPerSlice, offset)
	if err != nil {
		return nil, err
	}
	totalPages := utils.CalculateTotalPages(count, env.RecordsPerSlice)
	hasNextPage := totalPages > intPage
	slice := types.CustomerSlice{
		Records:     records,
		HasNextPage: hasNextPage,
	}
	return &slice, nil
}

func (s service) CreateCustomer(customer *types.Customer) error {
	if err := s.store.CreateCustomer(*customer); err != nil {
		return err
	}
	return nil
}

func (s service) ReadCustomer(id string) (*types.Customer, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	record, err := s.store.ReadCustomer(intId)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (s service) UpdateCustomer(customer *types.Customer) error {
	if err := s.store.UpdateCustomer(customer); err != nil {
		return err
	}
	return nil
}

func (s service) DeleteCustomer(id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if err := s.store.DeleteCustomer(intId); err != nil {
		return err
	}
	return nil
}
