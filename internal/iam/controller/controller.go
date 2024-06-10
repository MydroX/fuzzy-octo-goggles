// Package controller provides the implementation of the IAM controller.
package controller

import (
	"MydroX/project-v/internal/iam/usecases"
	"MydroX/project-v/pkg/logger"
	"MydroX/project-v/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	logger   *logger.Logger
	validate *validator.Validate
	usecases usecases.UsecasesInterface
	// repository repository.RepositoryInterface
}

// NewController is the interface for the controller.
func NewController(l *logger.Logger, v *validator.Validate, u usecases.UsecasesInterface) ControllerInterface {
	return &controller{
		logger:   l,
		validate: v,
		usecases: u,
	}
}

type createUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Role     string `json:"role"`
}

func (c *controller) CreateUser(ctx *gin.Context) {

	var request createUserRequest

	err := ctx.BindJSON(request)
	if err != nil {
		response.InvalidRequest(c.logger, ctx)
		return
	}

	err = c.validate.Struct(request)
	if err != nil {
		response.InvalidRequest(c.logger, ctx)
		return
	}

	err = c.usecases.Create(ctx, request.Username, request.Password, request.Email, request.Role)
	if err != nil {
		response.InternalError(c.logger, ctx)
		return
	}

	response.CreationSuccess(ctx, "user created")
}

func (c *controller) GetUser(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (c *controller) UpdateUser(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (c *controller) DeleteUser(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (c *controller) AuthenticateUser(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}
