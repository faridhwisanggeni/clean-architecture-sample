package domain

type Transaction struct {
	ID           string  `json:"id"`
	Trx_id       string  `json:"name"`
	Total_amount float64 `json:"price"`
	Item         []Item
}

type Item struct {
	ID       string  `json:"id"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
