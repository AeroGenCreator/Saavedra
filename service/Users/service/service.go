package service

import (
	"Saavedra/service/Users/store"
	"Saavedra/service/Users/types"
	"math"
)

type Service interface {
	NewUser(user *types.User) (*types.User, error)
	UpdateUser(user *types.User) (*types.User, error)
	ListUsers(page int) ([]*types.User, int, int, int, error)
	SelectUser(id int) (*types.User, error)
	DeleteUser(id int) error
}

type service struct {
	store store.Store
}

func New(store store.Store) Service {
	return service{store: store}
}

func divisionRounded(a, b int) int {

	if b == 0 {
		panic("Division by zero")
	}

	quotient := float64(a) / float64(b)

	res := math.Ceil(quotient)

	return int(res)
}

func (s service) NewUser(user *types.User) (*types.User, error) {

	user, err := s.store.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s service) UpdateUser(user *types.User) (*types.User, error) {
	user, err := s.store.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s service) ListUsers(page int) ([]*types.User, int, int, int, error) {

	var limit int
	var offset int

	offset = (page - 1) * 30
	limit = 30

	records, total_recs, err := s.store.SelectUserFrontendPage(limit, offset)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	total_pages := divisionRounded(total_recs, limit)

	return records, page, total_pages, total_recs, nil
}

func (s service) SelectUser(id int) (*types.User, error) {
	user, err := s.store.SelectUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s service) DeleteUser(id int) error {
	err := s.store.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
