// Package usecases provides the implementation of the IAM usecases.
package usecases

import (
	"MydroX/project-v/internal/iam/repository"
	"MydroX/project-v/pkg/logger"
	"context"
)

type usecases struct {
	logger     *logger.Logger
	repository repository.RepositoryInterface
}

// NewUsecases is creating an interface for all the usecases of the service.
func NewUsecases(l *logger.Logger, r repository.RepositoryInterface) UsecasesInterface {
	return &usecases{
		logger:     l,
		repository: r,
	}
}

func (u *usecases) Create(ctx context.Context, username, password, email, role string) error {
	if role == "" {
		role = "GUEST"
	}

	err := u.repository.CreateUser(username, email, password, role)

	return err
}

func (u *usecases) Get(ctx *context.Context) {
	panic("not implemented") // TODO: Implement
}

func (u *usecases) Update(ctx *context.Context) {
	panic("not implemented") // TODO: Implement
}

func (u *usecases) Delete(ctx *context.Context) {
	panic("not implemented") // TODO: Implement
}

func (u *usecases) Auth(ctx *context.Context) {
	panic("not implemented") // TODO: Implement
}
