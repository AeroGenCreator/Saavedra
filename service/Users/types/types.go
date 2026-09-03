package types

import "errors"

var ErrAdminProtection = errors.New("admin protection: standard admin user cannot be deleted")

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NoIdUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserStr struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsersSlice struct {
	Records     []*User `json:"records"`
	HasNextPage bool    `json:"hasNextPage"`
}
