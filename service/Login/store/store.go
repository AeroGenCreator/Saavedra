package store

import (
	"Saavedra/service/Login/types"
	utilities "Saavedra/utils"
	"database/sql"
)

// Contratos Del servicio
type Store interface {
	GetUserInfo(request *types.User) (*types.User, error)
	InsertUserSession(userInfo *types.AuthUserSession) error
	FetchUserAuthToken(user *types.User) (*types.AuthUserSession, error)
}

// Objeto "cursor" SQL -> Con sus contratos
type store struct {
	db *sql.DB
}

// Instancia de -> Objeto -> Contratos -> SQL
func New(db *sql.DB) Store {
	return &store{db: db}
}

// CONTRACT: QUERY TO USER INFO
func (s *store) GetUserInfo(request *types.User) (*types.User, error) {
	q := `SELECT id, name, email, password FROM users WHERE email = ?;`
	dbUser := types.User{}

	err := s.db.QueryRow(q, request.Email).Scan(&dbUser.Id, &dbUser.Name, &dbUser.Email, &dbUser.Password)
	if err == sql.ErrNoRows {
		dbUser.Name = ""
		dbUser.Password = ""
		return &dbUser, nil
	} else if err != nil {
		return nil, err
	}
	return &dbUser, nil
}

// CONTRATO: INSERT OR UPDATE USER SESSION IF NOT EXISTS
func (s *store) InsertUserSession(userInfo *types.AuthUserSession) error {

	query := `
	INSERT INTO auth_user_session (user_id, auth_token)
	VALUES (?, ?)
	ON CONFLICT(user_id) DO UPDATE SET auth_token = excluded.auth_token;
	`

	_, err := s.db.Exec(query, userInfo.UserId, userInfo.AuthToken)
	if err != nil {
		return err
	}

	return nil
}

// FUNCTION: CALL WHEN MAIN RUNS. INJECTS ADMIN CREDENTIALS.
func InjectAdminDB(admin types.User, db *sql.DB) error {

	hashedPassword, err := utilities.HashPassword(admin.Password)
	if err != nil {
		return err
	}

	admin.Password = hashedPassword

	command := `
	INSERT INTO users (name, email, password)
	VALUES (?, ?, ?)
	ON CONFLICT DO UPDATE SET
	name = excluded.name,
	email = excluded.email,
	password = excluded.password;
	`

	if _, err = db.Exec(command, admin.Name, admin.Email, admin.Password); err != nil {
		return err
	}

	return nil
}

// #
func (s *store) FetchUserAuthToken(user *types.User) (*types.AuthUserSession, error) {
	return nil, nil
}
