package types

import "errors"

var ErrUniqueConstraint = errors.New("Duplicated data. Skiping...")

// 3 TABLES FOR THIS SERVICE
// SUGGESTION: REAL FIELDS FIRST, RELATIONALS DENOTATE THEM WITH Id SUFFIX
// FOREIGN FIELDS LAST
type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	MaterialId  int     `json:"material"`
	ProveedorId int     `json:"proveedor"`
	Phone       string  `json:"phone"`
}

type Material struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Proveedor struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone int    `json:"number"`
}
