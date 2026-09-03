package types

import "errors"

var ErrNoRecord = errors.New("There is no record for the provided 'id'.")

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

// SLICES
type MaterialSlice struct {
	Records     []*Material `json:"records"`
	HasNextPage bool        `json:"hasNextPage"`
}

// ID STR
type MaterialStr struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
