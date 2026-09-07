package mCustomer

import (
	"database/sql"
	"fmt"
)

func CreateSchema(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// IF NEW DB: CREATION ELSE DATA MIGRATION (client)
	var exists bool
	checkCustomer := `SELECT EXISTS(SELECT 1 FROM sqlite_master WHERE type='table' AND name='customer');`
	if err := tx.QueryRow(checkCustomer).Scan(&exists); err != nil {
		return fmt.Errorf("Error validating customer table exists: %w", err)
	}

	if exists {
		qMigrate := `
		CREATE TABLE customer_new (
			id INTEGER PRIMARY KEY,
			name CHAR(20),
			fullname TEXT,
			address TEXT,
			technicianphone CHAR(10),
			buyerphone CHAR(10),
			customeremail TEXT,
			UNIQUE(name, fullname)
		);

		INSERT INTO customer_new (
			id, name, fullname, address,
			technicianphone, buyerphone, customeremail
		)
		SELECT
			id, name, fullname, address,
			technicianphone, buyerphone, customeremail
		FROM customer;

		DROP TABLE customer;
		ALTER TABLE customer_new RENAME TO customer;
		`

		if _, err := tx.Exec(qMigrate); err != nil {
			return fmt.Errorf("error migrando customer: %w", err)
		}
	} else {
		qCreate := `
		CREATE TABLE IF NOT EXISTS customer (
			id INTEGER PRIMARY KEY,
			name CHAR(20),
			fullname TEXT,
			address TEXT,
			technicianphone CHAR(10),
			buyerphone CHAR(10),
			customeremail TEXT,
			UNIQUE(name, fullname)
		);`

		if _, err := tx.Exec(qCreate); err != nil {
			return fmt.Errorf("Error creating customer: %w", err)
		}
	}

	return tx.Commit()
}
