package iam

import (
	"MydroX/project-v/internal/iam/controller"
	"MydroX/project-v/pkg/logger"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewServer(config *Config, logger *logger.Logger, validate *validator.Validate) {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	c := controller.NewController(logger, validate)

	// - Middleware SECRET KEY API for every endpoint in headers

	v1 := router.Group("/v1")
	v1.POST("/register", c.CreateUser)
	v1.POST("/auth", c.AuthenticateUser)
	v1.POST("/:uuid", c.GetUser)

	// TODO
	// - Middleware authentification
	// - UpdateUser
	// - DeleteUser

	err := router.Run(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		// logger.Fatal("error starting server", zap.Error(err))
	}
}
