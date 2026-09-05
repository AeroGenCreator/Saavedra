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

	// CHECK TABLE MATERIAL
	var material bool
	qm := `SELECT EXISTS(SELECT 1 FROM sqlite_master WHERE type='table' AND name='material');`
	err = tx.QueryRow(qm).Scan(&material)
	if err != nil {
		log.Printf("Table material exists error (%v)", err)
	}

	if material {
		migrateMaterial := `
		CREATE TABLE material_new (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE
		);
		INSERT INTO material_new(id, name)
		SELECT id, name FROM material;
		DROP TABLE material;
		ALTER table material_new RENAME TO material;
		`
		_, err := tx.Exec(migrateMaterial)
		if err != nil {
			log.Printf("Error migrating table material...(%v)", err.Error())
			return err
		}
	} else {
		createMaterial := `
		CREATE TABLE material (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE
		);
		`
		_, err := tx.Exec(createMaterial)
		if err != nil {
			log.Panicf("Error creating table material...(%v)", err.Error())
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
		material_id INTEGER,
		proveedor_id INTEGER,
		FOREIGN KEY (material_id) REFERENCES material(id) ON DELETE SET NULL,
		FOREIGN KEY (proveedor_id) REFERENCES proveedor(id) ON DELETE SET NULL
		);
		INSERT INTO product_new(id, name, description, pmeasure, price, material_id, proveedor_id)
		SELECT id, name, description, pmeasure, price, material_id, proveedor_id FROM product;
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
		material_id INTEGER,
		proveedor_id INTEGER,
		FOREIGN KEY (material_id) REFERENCES material(id) ON DELETE SET NULL,
		FOREIGN KEY (proveedor_id) REFERENCES proveedor(id) ON DELETE SET NULL
		);`
		_, err := tx.Exec(createProduct)
		if err != nil {
			log.Panicf("Error creating table product...(%v)", err.Error())
			return err
		}
	}

	return tx.Commit()
}
