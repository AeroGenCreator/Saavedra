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
	UpdateUserSession(authUser *types.AuthUserSession) error
	SelectUserSession(authUser *types.AuthUserSession) (*types.AuthUserSession, error)
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

// FUNCTION: CALL WHEN MAIN RUNS. INJECTS ADMIN CREDENTIALS.
func InjectAdminDB(admin types.User, db *sql.DB) (int, error) {

	hashedPassword, err := utilities.HashPassword(admin.Password)
	if err != nil {
		return 0, err
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

	res, err := db.Exec(command, admin.Name, admin.Email, admin.Password)
	if err != nil {
		return 0, err
	}

	adminId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(adminId), nil
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

// CONTRACT: WHEN LOG OUT -> NO JWT TOKEN || UPDATE
func (s *store) UpdateUserSession(authUser *types.AuthUserSession) error {
	q := "UPDATE auth_user_session SET auth_token = ? WHERE user_id = ?;"

	_, err := s.db.Exec(q, authUser.AuthToken, authUser.UserId)
	if err != nil {
		return err
	}
	return nil
}

// CONTRACT: FETCH DB APY TOKEN
func (s *store) SelectUserSession(authUser *types.AuthUserSession) (*types.AuthUserSession, error) {
	q := "SELECT id, user_id, auth_token FROM auth_user_session WHERE user_id = ?;"
	session := types.AuthUserSession{}
	err := s.db.QueryRow(q, authUser.UserId).Scan(&session.Id, &session.UserId, &session.AuthToken)
	if err == sql.ErrNoRows {
		session.AuthToken = ""
		return &session, nil
	} else if err != nil {
		return nil, err
	}
	return &session, nil
}
