package response

import (
	"MydroX/project-v/pkg/logger"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Error is a function to handle error response
// The function will log the error message and return the error message to the client.
// In the case of debug mode, the error message will be shown, otherwise, a generic error message will be shown.
func Error(logger *logger.Logger, ctx *gin.Context, code int, message string) {
	logger.Zap.Error(fmt.Sprintf("[%d] %s", code, message))
	if logger.Debug {
		ctx.JSON(code, gin.H{"error": message})
		return
	}
	ctx.JSON(code, gin.H{"error": "an error occurred. please try again later or contact support"})
}

// CreationSuccess is a function to handle success response for creation of any entity
func CreationSuccess(ctx *gin.Context, message string) {
	ctx.JSON(201, gin.H{"message": message})
	return
}
