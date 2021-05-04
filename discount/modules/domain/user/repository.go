package user

type Repository interface {
	GetById(id string) (User, error)
}
