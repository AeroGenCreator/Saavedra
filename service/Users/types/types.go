package types

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Rows struct {
	Records []*User `json:"records"`
}

type Metadata struct {
	CountRows   int `json:"count_rows"`
	CountPages  int `json:"count_pages"`
	CurrentPage int `json:"current_page"`
}

type Table struct {
	Data *Rows     `json:"data"`
	Info *Metadata `json:"info"`
}
