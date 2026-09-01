package types

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Table struct {
	Rows        []*User `json:"rows"`
	CountRows   int     `json:"count_rows"`
	CountPages  int     `json:"count_pages"`
	CurrentPage int     `json:"current_page"`
}
