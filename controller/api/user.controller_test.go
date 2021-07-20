package api_test

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"platform-sample/controller/api"
	"platform-sample/infrastructure/database"
	"platform-sample/infrastructure/server"
	mocks2 "platform-sample/mocks/service"
	"platform-sample/model"
	"platform-sample/service"
	"strings"
	"testing"
)

func initIntegrateMockUserService() *service.UserServiceImpl {
	mockDb := database.SqlStore{}.GetMockDb()
	mockServer := server.Server{MainDb: mockDb}
	return mockServer.InjectUserService()
}

func initUnitMockUserService() {
}

func Test_IntegrateCreateUser(t *testing.T) {
	mockUser := model.User{Name: "James"}
	byteData, err := json.Marshal(mockUser)
	if err != nil {
		log.Println(err)
	}

	// 두개 뭔 차이지..?
	// bytes.NewReader()
	// bytes.NewBuffer()
	buff := bytes.NewBuffer(byteData)

	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/users", buff)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	userController := api.UserController{}.NewUserController(initIntegrateMockUserService())
	if assert.NoError(t, userController.CreateUser(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), mockUser.Name)
	}
}

func Test_Unit_CreateUser(t *testing.T) {
	mockUser := &model.User{Name: "James"}
	byteData, err := json.Marshal(mockUser)
	if err != nil {
		log.Println(err)
	}

	buff := bytes.NewBuffer(byteData)

	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/users", buff)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)

	mockService := &mocks2.UserService{}
	mockService.On("CreateUser", mockUser).Return(mockUser, nil)
	userController := api.UserController{}.NewUserController(mockService)

	if assert.NoError(t, userController.CreateUser(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		//assert.Contains(t, rec.Body.String(), mockUser.Name)
		assert.Equal(t, string(byteData), strings.Trim(rec.Body.String(), "\n"))
	}
}
