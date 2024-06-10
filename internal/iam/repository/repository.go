package repository

import (
	"MydroX/project-v/internal/models"
	"MydroX/project-v/pkg/logger"

	"gorm.io/gorm"
)

type repository struct {
	logger *logger.Logger
	db     *gorm.DB
}

func NewRepository(l *logger.Logger, db *gorm.DB) RepositoryInterface {
	return &repository{
		logger: l,
		db:     db,
	}
}

func (r *repository) CreateUser(username, email, password, role string) error {
	user := &models.User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}

	res := r.db.Create(user)
	if res.Error != nil {
		r.logger.Zap.Sugar().Errorf("error creating user: %v", res.Error)
		return res.Error
	}

	return nil
}
