package mProduct

import (
	"database/sql"
	"log"
)

func CreateSchema(db *sql.DB) error {

	// START SQL TRANSACTION
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	// CHECK TABLE CATEGORY
	var category bool
	qm := `SELECT EXISTS(SELECT 1 FROM sqlite_master WHERE type='table' AND name='category');`
	err = tx.QueryRow(qm).Scan(&category)
	if err != nil {
		log.Printf("Table category exists error (%v)", err)
	}

	if category {
		migrateCategory := `
		CREATE TABLE category_new (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE
		);
		INSERT INTO category_new(id, name)
		SELECT id, name FROM category;
		DROP TABLE category;
		ALTER table category_new RENAME TO category;
		`
		_, err := tx.Exec(migrateCategory)
		if err != nil {
			log.Printf("Error migrating table category...(%v)", err.Error())
			return err
		}
	} else {
		createCategory := `
		CREATE TABLE category (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE
		);
		`
		_, err := tx.Exec(createCategory)
		if err != nil {
			log.Panicf("Error creating table category...(%v)", err.Error())
			return err
		}
	}

	// CHECK TABLE PROVEEDOR
	var proveedor bool
	qp := `SELECT EXISTS(SELECT 1 FROM sqlite_master WHERE type='table' AND name='proveedor');`
	err = tx.QueryRow(qp).Scan(&proveedor)
	if err != nil {
		log.Printf("Table proveedor exists error (%v)", err)
	}

	if proveedor {
		migrateProveedor := `
		CREATE TABLE proveedor_new (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		phone INTEGER
		);
		INSERT INTO proveedor_new(id, name, phone)
		SELECT id, name, phone FROM proveedor;
		DROP TABLE proveedor;
		ALTER table proveedor_new RENAME TO proveedor;
		`
		_, err := tx.Exec(migrateProveedor)
		if err != nil {
			log.Printf("Error migrating table proveedor...(%v)", err.Error())
			return err
		}
	} else {
		createProveedor := `
		CREATE TABLE proveedor (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		phone INTEGER
		);
		`
		_, err := tx.Exec(createProveedor)
		if err != nil {
			log.Panicf("Error creating table proveedor...(%v)", err.Error())
			return err
		}
	}

	// CHECK TABLE PRODUCT
	var product bool
	qprod := `SELECT EXISTS(SELECT 1 FROM sqlite_master WHERE type='table' AND name='product');`
	err = tx.QueryRow(qprod).Scan(&product)
	if err != nil {
		log.Printf("Table product exists error (%v)", err)
	}

	if product {
		migrateProduct := `
		CREATE TABLE product_new (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		pmeasure TEXT NOT NULL,
		price FLOAT NOT NULL,
		category_id INTEGER,
		proveedor_id INTEGER,
		FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE SET NULL,
		FOREIGN KEY (proveedor_id) REFERENCES proveedor(id) ON DELETE SET NULL,
		UNIQUE(name, pmeasure, proveedor_id)
		);
		INSERT INTO product_new(id, name, description, pmeasure, price, category_id, proveedor_id)
		SELECT id, name, description, pmeasure, price, category_id, proveedor_id FROM product;
		DROP TABLE product;
		ALTER table product_new RENAME TO product;
		`
		_, err := tx.Exec(migrateProduct)
		if err != nil {
			log.Printf("Error migrating table product...(%v)", err.Error())
			return err
		}
	} else {
		createProduct := `
		CREATE TABLE product (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		pmeasure TEXT NOT NULL,
		price FLOAT NOT NULL,
		category_id INTEGER,
		proveedor_id INTEGER,
		FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE SET NULL,
		FOREIGN KEY (proveedor_id) REFERENCES proveedor(id) ON DELETE SET NULL,
		UNIQUE(name, pmeasure, proveedor_id)
		);`
		_, err := tx.Exec(createProduct)
		if err != nil {
			log.Panicf("Error creating table product...(%v)", err.Error())
			return err
		}
	}

	return tx.Commit()
}
