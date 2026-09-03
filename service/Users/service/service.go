package service

import (
	"Saavedra/env"
	"Saavedra/service/Users/store"
	"Saavedra/service/Users/types"
	"Saavedra/utils"
	"math"
	"strconv"
)

type Service interface {
	NewUser(user *types.User) (*types.User, error)
	UpdateUser(user *types.User) (*types.User, error)
	ListUsers(page string) (*types.UsersSlice, error)
	SelectUser(id int) (*types.User, error)
	DeleteUser(user *types.User) error
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

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	newUser, err := s.store.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s service) UpdateUser(user *types.User) (*types.User, error) {

	if user.Password == "" {
		user, err := s.store.UpdateUserNoPass(user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	user, err = s.store.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// RETURNS: rows, current_page, total_records, total_pages, error
func (s service) ListUsers(page string) (*types.UsersSlice, error) {

	intPage, err := strconv.Atoi(page)
	if err != nil {
		intPage = 1
	}

	var offset int
	offset = (intPage - 1) * env.RecordsPerSlice
	records, count, err := s.store.SelectAllUsers(env.RecordsPerSlice, offset)
	if err != nil {
		return nil, err
	}

	totalPages := utils.CalculateTotalPages(count, env.RecordsPerSlice)
	hasNextPage := totalPages > intPage
	userSlice := types.UsersSlice{
		Records:     records,
		HasNextPage: hasNextPage,
	}

	return &userSlice, nil
}

func (s service) SelectUser(id int) (*types.User, error) {
	user, err := s.store.SelectUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s service) DeleteUser(user *types.User) error {
	if user.Id == env.AdminId {
		return types.ErrAdminProtection
	}
	err := s.store.DeleteUser(user)
	if err != nil {
		return err
	}

	return nil
}
