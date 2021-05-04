package product

type Repository interface {
	GetAll() ([]Product, error)
}
