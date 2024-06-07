package rest

import (
	"MydroX/project-v/internal/iam"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewServer(config *iam.Config, logger zap.Logger) {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	err := router.Run(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		logger.Fatal("error starting server", zap.Error(err))
	}
}
