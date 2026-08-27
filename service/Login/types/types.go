package types

import "github.com/golang-jwt/jwt/v4"

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthUserSession struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	AuthToken string `json:"auth_token"`
}

type Credentials struct {
	UserName         string `json:"user_name"`
	ValidationStatus bool   `json:"validation_status"`
	Token            string `json:"token"`
}

type Claims struct {
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

type CompanyBranding struct {
	CompanyName string `json:"company_name"`
	Slogan      string `json:"slogan"`
}
