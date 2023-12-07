package domain

type Transaksi struct {
	Id         string
	ProductId  string
	Quantity   int
	OrderId    string
	TotalPrice float64
}
