package controller

import "github.com/gin-gonic/gin"

type ControllerInterface interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	AuthenticateUser(c *gin.Context)
}
