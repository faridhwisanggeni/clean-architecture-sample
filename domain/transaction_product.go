package domain

type TransactionProduct struct {
	ID                     int                 `json:"transaction_id,omitempty"`
	TransactionCreatedDate string              `json:"transaction_created_date,omitempty"`
	ProductPriceTotal      float64             `json:"product_price_total"`
	ProductQtyTotal        int                 `json:"product_qty_total"`
	ProductDetail          []TransactionDetail `json:"trx_detail"`
}
