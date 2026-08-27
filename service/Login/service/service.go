package service

import (
	store "Saavedra/service/Login/store"
	"Saavedra/service/Login/types"
	utils "Saavedra/utils"
	"encoding/json"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	ValidatePassword(req *types.LoginRequest) (*types.Credentials, error)
	LoadCompanyBranding() (*types.CompanyBranding, error)
}

type service struct {
	store store.Store
}

func New(store store.Store) Service {
	return service{store: store}
}

// FUNCIÓN: CHECK PASSWORD
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CONTRACT POST: VALIDATE LOGIN
func (s service) ValidatePassword(req *types.LoginRequest) (*types.Credentials, error) {

	// 1. QUERY USER
	dbNamePassword, err := s.store.GetUserInfo(req)
	if err != nil {
		return nil, err
	}

	var creds types.Credentials
	var userSession types.AuthUserSession
	var token = ""

	// 2. VALIDATE LOGIN
	boolean := CheckPasswordHash(req.Password, dbNamePassword.HashedPassword)

	// 3. GENERATE TOKEN
	if boolean && dbNamePassword.UserName != "" {
		token, err = utils.GenerateToken(dbNamePassword.UserName)
		if err != nil {
			return nil, err
		}
	}

	// 4. NEW USER? INSERT AUTH IN DB
	userSession.UserId = dbNamePassword.Id
	userSession.AuthToken = token

	err = s.store.InsertUserSession(&userSession)
	if err != nil {
		return nil, err
	}

	// 5. RETUNR CREDENTIALS
	creds.UserName = dbNamePassword.UserName
	creds.ValidationStatus = boolean
	creds.Token = token

	return &creds, nil
}

// CONTRATO GET Exportar Branding
func (s service) LoadCompanyBranding() (*types.CompanyBranding, error) {
	// Lectura en bytes del fichero
	fileBytes, err := os.ReadFile("config/global.json")
	if err != nil {
		return nil, err
	}
	// Parsing y asignacion a un "type"
	var brand types.CompanyBranding
	if err := json.Unmarshal(fileBytes, &brand); err != nil {
		return nil, err
	}
	return &brand, nil
}
