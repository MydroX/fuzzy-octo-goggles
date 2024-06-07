package rest

import (
	"MydroX/project-v/internal/iam"

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

	router.Run()
}
