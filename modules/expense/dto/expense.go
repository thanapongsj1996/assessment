package dto

type AddExpenseReq struct {
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

type AddExpenseRes struct {
	ID int `json:"id"`
	AddExpenseReq
}
