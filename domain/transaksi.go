package domain

type Transaksi struct {
	ID         int               `json:"id"`
	TotalQty   int               `json:"total_qty"`
	TotalPrice float64           `json:"total_price"`
	Item       []TransaksiDetail `json:"item"`
}

type TransaksiDetail struct {
	ID          string  `json:"id"`
	IDTransaksi int     `json:"id_transaksi"`
	IDProduct   string  `json:"id_product"`
	Qty         int     `json:"qty"`
	BasicPrice  float64 `json:"basic_price"`
	Total       float64 `json:"total"`
}
