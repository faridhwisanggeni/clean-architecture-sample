package domain

type TransactionDetail struct {
	ID            int     `json:"id,omitempty"`
	TransactionID int     `json:"transaction_id,omitempty"`
	ProductID     int     `json:"product_id"`
	ProductPrice  float64 `json:"product_price"`
	ProductQty    int     `json:"product_qty"`
}
