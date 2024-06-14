// Package usecases provides the implementation of the IAM usecases.
package usecases

import (
	"MydroX/project-v/internal/iam/users/dto"
	"MydroX/project-v/internal/iam/users/models"
	"MydroX/project-v/internal/iam/users/repository"
	"MydroX/project-v/pkg/logger"
	"MydroX/project-v/pkg/password"
	"fmt"

	"github.com/google/uuid"
)

type usecases struct {
	logger     *logger.Logger
	repository repository.UsersRepository
}

// NewUsecases is creating an interface for all the usecases of the service.
func NewUsecases(l *logger.Logger, r repository.UsersRepository) UsersUsecases {
	return &usecases{
		logger:     l,
		repository: r,
	}
}

func (u *usecases) Create(req dto.CreateUserRequest) error {
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Role:     req.Role,
	}

	passwordCrypted, err := password.Hash(req.Password)
	if err != nil {
		return err
	}
	user.Password = passwordCrypted

	userUUID := uuid.New()
	user.UUID = userUUID

	err = u.repository.CreateUser(user)

	return err
}

func (u *usecases) Get(uuid uuid.UUID) (*dto.GetUserResponse, error) {
	user, err := u.repository.GetUser(uuid)

	res := dto.GetUserResponse{
		UUID:     user.UUID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}

	return &res, err
}

func (u *usecases) Update(user models.User) error {
	err := u.repository.UpdateUser(user)
	return err
}

func (u *usecases) Delete(uuid uuid.UUID) error {
	err := u.repository.DeleteUser(uuid)
	return err
}

func (u *usecases) Auth(username, email, password string) (string, error) {
	return "", fmt.Errorf("not implemented")
}
