package repository

type RepositoryInterface interface {
	CreateUser(name, email, password, role string) error
}
