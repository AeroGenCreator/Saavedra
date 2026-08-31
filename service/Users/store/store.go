package store

import (
	"Saavedra/service/Users/types"
	"database/sql"
)

type Store interface {
	CreateUser(user *types.User) (*types.User, error)
	UpdateUser(user *types.User) (*types.User, error)
	SelectUserFrontendPage(limit int, offset int) ([]*types.User, int, error)
	SelectUser(id int) (*types.User, error)
	DeleteUser(id int) error
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return &store{db: db}
}

func (s *store) CreateUser(user *types.User) (*types.User, error) {
	q := "INSERT INTO users (name, email, password) VALUES (?, ?, ?);"

	res, err := s.db.Exec(q, user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.Id = int(id)

	return user, nil
}

func (s *store) UpdateUser(user *types.User) (*types.User, error) {
	q := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?;"

	_, err := s.db.Exec(q, user.Name, user.Email, user.Password, user.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *store) SelectUserFrontendPage(limit int, offset int) ([]*types.User, int, error) {

	q1 := "SELECT COUNT(id) AS count_id FROM users GROUP BY id;"
	q2 := "SELECT id, name, email, password FROM users LIMIT ? OFFSET ?;"

	var total_recs int
	err := s.db.QueryRow(q1).Scan(&total_recs)
	if err != nil {
		return nil, 0, err
	}

	rows, err := s.db.Query(q2, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var records []*types.User

	for rows.Next() {
		var user types.User

		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, 0, err
		}
		records = append(records, &user)
	}

	if rows.Err() != nil {
		return nil, 0, rows.Err()
	}

	return records, total_recs, nil
}

func (s *store) SelectUser(id int) (*types.User, error) {
	q := "SELECT id, name, email, password FROM users WHERE id = ?;"

	var user types.User
	err := s.db.QueryRow(q, id).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *store) DeleteUser(id int) error {
	q := "DELETE FROM users WHERE id = ?;"

	if _, err := s.db.Exec(q, id); err != nil {
		return err
	}

	return nil
}
