package domain

type ProductUseCase interface {
	GetProducts() ([]*Product, error)
	GetProductByID(id string) (*Product, error)
	CreateProduct(product *Product) (*Product, error)
	UpdateProduct(id string, name string, price float64) (*Product, error)
	DeleteProduct(id string) error
}
