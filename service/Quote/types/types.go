package types

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Customer struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	CustomerEmail string `json:"customerEmail"`
}

type ProductFetch struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	PMeasure    string  `json:"pMeasure"`
	Price       float32 `json:"price"`
	CategoryId  int     `json:"categoryId"`
	Category    string  `json:"category"`
	ProveedorId int     `json:"proveedorId"`
	Proveedor   string  `json:"proveedor"`
}

type Many2One struct {
	UserArray     []*User         `json:"userArray"`
	CustomerArray []*Customer     `json:"customerArray"`
	ProductArray  []*ProductFetch `json:"productArray"`
}
