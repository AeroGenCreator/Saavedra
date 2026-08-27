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

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DBUserInfo struct {
	Id             int    `json:"id"`
	UserName       string `json:"user_name"`
	HashedPassword string `json:"hashed_password"`
}

type Credentials struct {
	UserName         string `json:"user_name"`
	ValidationStatus bool   `json:"validation_status"`
	Token            string `json:"token"`
}

// Struct: Signed Token
type Claims struct {
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

type CompanyBranding struct {
	CompanyName string `json:"company_name"`
	Slogan      string `json:"slogan"`
	IconPath    string `json:"icon_path"`
}
