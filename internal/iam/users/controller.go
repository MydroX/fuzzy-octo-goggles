package users

import (
	"MydroX/project-v/internal/iam/users/dto"
	"MydroX/project-v/internal/iam/users/usecases"
	apiError "MydroX/project-v/pkg/errors"
	"MydroX/project-v/pkg/logger"
	"MydroX/project-v/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Controller struct {
	logger   *logger.Logger
	validate *validator.Validate
	usecases usecases.UsersUsecases
}

// NewController is the interface for the controller.
func NewController(l *logger.Logger, u usecases.UsersUsecases) *Controller {
	validator := validator.New()

	return &Controller{
		validate: validator,
		logger:   l,
		usecases: u,
	}
}

func (c *Controller) CreateUser(ctx *gin.Context) {
	var request dto.CreateUserRequest

	err := ctx.BindJSON(&request)
	if err != nil {
		response.InvalidRequest(c.logger, ctx)
		return
	}

	err = c.validate.Struct(request)
	if err != nil {
		response.InvalidRequest(c.logger, ctx)
		return
	}

	// ^[A-Za-z0-9._-]{4,18}$  // username

	err = c.usecases.Create(request)
	if err != nil {
		response.InternalError(c.logger, ctx, err)
		return
	}

	response.CreationSuccess(ctx, "user created")
}

func (c *Controller) GetUser(ctx *gin.Context) {
	uuidStr := ctx.Param("uuid")
	if uuidStr == "" {
		response.InvalidRequest(c.logger, ctx)
		return
	}

	err := uuid.Validate(uuidStr)
	if err != nil {
		response.InvalidRequest(c.logger, ctx)
		return
	}

	userUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		response.InvalidRequest(c.logger, ctx)
		return
	}

	resp, err := c.usecases.Get(userUUID)
	if err != nil {
		if err == apiError.ErrNotFound {
			response.NotFound(c.logger, ctx)
			return
		}
		response.InternalError(c.logger, ctx, err)
		return
	}

	ctx.JSON(200, resp)
}

func (c *Controller) UpdateUser(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
}

func (c *Controller) DeleteUser(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
}

func (c *Controller) AuthenticateUser(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
}
