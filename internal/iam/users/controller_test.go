package users

import (
	"MydroX/project-v/internal/iam/users/dto"
	"MydroX/project-v/internal/iam/users/mocks"
	"MydroX/project-v/pkg/logger"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	v1 = "/api/v1"

	create = "/register"
)

type testServer struct {
	router      *gin.Engine
	mockUsecase *mocks.MockUsersUsecases
}

func testRouter(logger *logger.Logger, _ *gorm.DB, controller Controller) *gin.Engine {
	router := gin.Default()

	err := router.SetTrustedProxies(nil)
	if err != nil {
		logger.Zap.Fatal("error setting trusted proxies", zap.Error(err))
	}

	api := router.Group("api")
	v1 := api.Group("/v1")

	users := v1.Group("/users")
	users.POST("/register", controller.CreateUser)
	users.POST("/auth", controller.AuthenticateUser)
	users.POST("/:uuid", controller.GetUser)

	// TODO: Middleware authentification
	users.POST("/update/:uuid", controller.UpdateUser)
	users.DELETE("/:uuid", controller.DeleteUser)

	return router
}

func newServerTest(t *testing.T) testServer {
	logger := logger.New("TEST")

	ctrl := gomock.NewController(t)
	usecasesMock := mocks.NewMockUsersUsecases(ctrl)

	c := NewController(logger, usecasesMock)
	router := testRouter(logger, nil, *c)

	return testServer{
		router:      router,
		mockUsecase: usecasesMock,
	}
}

func Test_Create(t *testing.T) {
	s := newServerTest(t)

	t.Run("Create with success V1", func(t *testing.T) {
		input := dto.CreateUserRequest{
			Username: "test",
			Email:    "test@test.com",
			Password: "thisisatestpassword1234!@#$",
			Role:     "USER",
		}
		userJSON, _ := json.Marshal(input)

		req, _ := http.NewRequest("POST", v1+create, strings.NewReader(string(userJSON)))

		s.mockUsecase.EXPECT().Create(gomock.Any()).Return(nil)

		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("Failed to bind JSON", func(t *testing.T) {
		req, _ := http.NewRequest("POST", v1+create, strings.NewReader(""))

		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Failed to validate JSON", func(t *testing.T) {
		input := dto.CreateUserRequest{
			Username: "test",
			Email:    "",
			Password: "thisisatestpassword1234!@#$",
			Role:     "USER",
		}
		userJSON, _ := json.Marshal(input)

		req, _ := http.NewRequest("POST", v1+create, strings.NewReader(string(userJSON)))

		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Usecase error V1", func(t *testing.T) {
		input := dto.CreateUserRequest{
			Username: "test",
			Email:    "test@test.com",
			Password: "thisisatestpassword1234!@#$",
			Role:     "USER",
		}
		userJSON, _ := json.Marshal(input)

		req, _ := http.NewRequest("POST", v1+create, strings.NewReader(string(userJSON)))

		s.mockUsecase.EXPECT().Create(gomock.Any()).Return(fmt.Errorf("test error"))

		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
