package repository

import "github.com/google/uuid"

type RepositoryInterface interface {
	CreateUser(uuid uuid.UUID, name, email, password, role string) error
}
