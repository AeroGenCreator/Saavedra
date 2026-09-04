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
	PMeasure    string  `json:"pMeasure"`
	Price       float32 `json:"price"`
	MaterialId  int     `json:"materialId"`
	ProveedorId int     `json:"proveedorId"`
	Phone       string  `json:"phone"`
}

type PMeasure struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
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

type ProveedorSlice struct {
	Records     []*Proveedor `json:"records"`
	HasNextPage bool         `json:"hasNextPage"`
}

type PMeasureSlice struct {
	Records []*PMeasure `json:"records"`
}

// ID STR
type MaterialStr struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ProveedorStr struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
