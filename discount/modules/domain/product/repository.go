package product

type Repository interface {
	GetById(id string) (Product, error)
	GetAll() ([]Product, error)
}
