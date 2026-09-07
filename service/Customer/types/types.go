package types

type Customer struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	FullName        string `json:"fullName"`
	Address         string `json:"address"`
	TechnicianPhone string `json:"technicianPhone"`
	BuyerPhone      string `json:"buyerPhone"`
	CustomerEmail   string `json:"customerEmail"`
}

type CustomerStr struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	FullName        string `json:"fullName"`
	Address         string `json:"address"`
	TechnicianPhone string `json:"technicianPhone"`
	BuyerPhone      string `json:"buyerPhone"`
	CustomerEmail   string `json:"customerEmail"`
}

type CustomerSlice struct {
	Records     []*Customer `json:"records"`
	HasNextPage bool        `json:"hasNextPage"`
}
