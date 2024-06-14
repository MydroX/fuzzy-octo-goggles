package usecases

import (
	"MydroX/project-v/internal/iam/users/dto"
	"MydroX/project-v/internal/iam/users/models"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=../mocks/mock_usecases.go -imports=models=MydroX/project-v/internal/users/models -package=mocks MydroX/project-v/internal/iam/users/usecases UsersUsecases

// Usecases is the interface to all the implemented usecases
type UsersUsecases interface {
	Create(user dto.CreateUserRequest) error
	Get(uuid uuid.UUID) (*dto.GetUserResponse, error)
	Update(user models.User) error
	Delete(uuid uuid.UUID) error
	Auth(username, email, password string) (string, error)
}
