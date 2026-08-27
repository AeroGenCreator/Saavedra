package mLogin

import (
	"database/sql"
	"fmt"
)

func CreateSchema(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// BACKUP FOR THE PROCESS
	defer tx.Rollback()

	// IF NEW DB: CREATION ELSE DATA MIGRATION (users)
	var exists bool
	queryCheckUsers := `SELECT EXISTS(SELECT 1 FROM sqlite_master WHERE type='table' AND name='users');`
	if err := tx.QueryRow(queryCheckUsers).Scan(&exists); err != nil {
		return fmt.Errorf("error verificando existencia de users: %w", err)
	}

	if exists {
		q1Migrate := `
		CREATE TABLE users_new (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
		INSERT INTO users_new (id, name, email, password)
		SELECT id, name, email, password FROM users;
		DROP TABLE users;
		ALTER TABLE users_new RENAME TO users;
		`
		if _, err := tx.Exec(q1Migrate); err != nil {
			return fmt.Errorf("error migrando users: %w", err)
		}
	} else {
		q1Create := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);`
		if _, err := tx.Exec(q1Create); err != nil {
			return fmt.Errorf("error creando users: %w", err)
		}
	}

	// IF NEW DB: CREATION ELSE DATA MIGRATION (auth_user_session)
	var sessionExists bool
	queryCheckSession := `SELECT EXISTS(SELECT 1 FROM sqlite_master WHERE type='table' AND name='auth_user_session');`
	if err := tx.QueryRow(queryCheckSession).Scan(&sessionExists); err != nil {
		return fmt.Errorf("error verificando existencia de auth_user_session: %w", err)
	}

	if sessionExists {
		q2Migrate := `
		CREATE TABLE auth_user_session_new (
			id INTEGER PRIMARY KEY,
			user_id INTEGER NOT NULL UNIQUE,
			auth_token TEXT,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);
		INSERT INTO auth_user_session_new (id, user_id, auth_token)
		SELECT id, user_id, auth_token FROM auth_user_session;
		DROP TABLE auth_user_session;
		ALTER TABLE auth_user_session_new RENAME TO auth_user_session;
		`
		if _, err := tx.Exec(q2Migrate); err != nil {
			return fmt.Errorf("error migrando auth_user_session: %w", err)
		}
	} else {
		q2Create := `
		CREATE TABLE IF NOT EXISTS auth_user_session (
			id INTEGER PRIMARY KEY,
			user_id INTEGER NOT NULL UNIQUE,
			auth_token TEXT,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);`
		if _, err := tx.Exec(q2Create); err != nil {
			return fmt.Errorf("error creando auth_user_session: %w", err)
		}
	}

	return tx.Commit()
}
