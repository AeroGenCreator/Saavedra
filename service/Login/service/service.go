package service

import (
	store "Saavedra/service/Login/store"
	"Saavedra/service/Login/types"
	utils "Saavedra/utils"
	"encoding/json"
	"os"
)

type Service interface {
	ValidatePassword(req *types.User) (*types.Credentials, error)
	LoadCompanyBranding() (*types.CompanyBranding, error)
}

type service struct {
	store store.Store
}

func New(store store.Store) Service {
	return service{store: store}
}

// CONTRACT POST: VALIDATE LOGIN
func (s service) ValidatePassword(req *types.User) (*types.Credentials, error) {

	// 1. QUERY USER
	dbUser, err := s.store.GetUserInfo(req)
	if err != nil {
		return nil, err
	}

	var creds types.Credentials
	var userSession types.AuthUserSession
	var token = ""

	// 2. VALIDATE LOGIN
	boolean := utils.CheckPasswordHash(req.Password, dbUser.Password)

	// 3. GENERATE TOKEN
	if boolean && dbUser.Name != "" {
		token, err = utils.GenerateToken(dbUser)
		if err != nil {
			return nil, err
		}
	}

	// 4. NEW USER? INSERT AUTH IN DB
	userSession.UserId = dbUser.Id
	userSession.AuthToken = token

	err = s.store.InsertUserSession(&userSession)
	if err != nil {
		return nil, err
	}

	// 5. RETUNR CREDENTIALS
	creds.UserName = dbUser.Name
	creds.ValidationStatus = boolean
	creds.Token = token

	return &creds, nil
}

// CONTRACT GET Exportar Branding
func (s service) LoadCompanyBranding() (*types.CompanyBranding, error) {

	// Lectura en bytes del fichero
	fileBytes, err := os.ReadFile("config/brand.json")
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
