package service

import (
	"Saavedra/service/Quote/store"
	"Saavedra/service/Quote/types"
)

type Service interface {
	Many2One() (*types.Many2One, error)
}

type service struct {
	store store.Store
}

func New(store store.Store) Service {
	return service{store: store}
}

func (s service) Many2One() (*types.Many2One, error) {
	many2one, err := s.store.Many2One()
	if err != nil {
		return nil, err
	}
	return many2one, nil
}
