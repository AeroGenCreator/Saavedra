package store

import (
	"Saavedra/service/Login/types"
	utilities "Saavedra/utils"
	"database/sql"
)

// Contratos Del servicio
type Store interface {
	GetUserInfo(request *types.LoginRequest) (*types.DBUserInfo, error)
	InsertUserSession(userInfo *types.AuthUserSession) error
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
func (s *store) GetUserInfo(request *types.LoginRequest) (*types.DBUserInfo, error) {
	q := `SELECT name, password FROM users WHERE email = ?;`
	dbUser := types.DBUserInfo{}

	err := s.db.QueryRow(q, request.Email).Scan(&dbUser.UserName, &dbUser.HashedPassword)
	if err == sql.ErrNoRows {
		dbUser.UserName = ""
		dbUser.HashedPassword = ""
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
	INSERT INTO users (name, email, password) VALUES (?, ?, ?)
	ON CONFLICT(email) DO UPDATE SET
	email = excluded.email,
	password = excluded.password;
	`

	if _, err = db.Exec(command, admin.Name, admin.Email, admin.Password); err != nil {
		return err
	}

	return nil
}
