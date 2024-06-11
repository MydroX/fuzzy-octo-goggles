// Package repository provides every implementation of database operations.
package repository

import (
	"MydroX/project-v/internal/models"
	"MydroX/project-v/pkg/logger"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	logger *logger.Logger
	db     *gorm.DB
}

// NewRepository is creating an interface for every method of the repository
func NewRepository(l *logger.Logger, db *gorm.DB) RepositoryInterface {
	return &repository{
		logger: l,
		db:     db,
	}
}

func (r *repository) CreateUser(uuid uuid.UUID, username, email, password, role string) error {
	user := &models.User{
		UUID:     uuid,
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}

	res := r.db.Create(user)
	if res.Error != nil {
		if res.Error == gorm.ErrDuplicatedKey {
			fmt.Println(res.Error)
		}
		r.logger.Zap.Sugar().Errorf("error creating user: %v", res.Error)
		return res.Error
	}

	return nil
}
