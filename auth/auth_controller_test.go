package auth

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/triaton/forum-backend-echo/common"
	"github.com/triaton/forum-backend-echo/common/utils"
	MocksUtils "github.com/triaton/forum-backend-echo/mocks/common/utils"
	MocksUsers "github.com/triaton/forum-backend-echo/mocks/users"
	"github.com/triaton/forum-backend-echo/test"
	"github.com/triaton/forum-backend-echo/users"
	UserModels "github.com/triaton/forum-backend-echo/users/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testName = "test"
var testEmail = "test-gin-boilerplate@test.com"
var testPassword = "123456"

func TestLoginFailWithParameterValidation(t *testing.T) {
	test.InitTest()
	testServer := echo.New()
	authController := AuthController{}
	var loginForm LoginRequest
	loginForm.Email = "non-email-format"
	loginForm.Password = "password"
	data, _ := json.Marshal(loginForm)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(string(data)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp := httptest.NewRecorder()
	context := testServer.NewContext(req, resp)

	if assert.NoError(t, authController.Login(context)) {
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	}
}

func TestLoginFailWithNonExistingUser(t *testing.T) {
	testServer := echo.New()
	testServer.Validator = &common.CustomValidator{Validator: validator.New()}
	authController := AuthController{}
	loginForm := LoginRequest{
		Email:    testEmail,
		Password: testPassword,
	}
	mockUserService := &MocksUsers.UsersService{}
	mockUserService.On("FindUserByEmail", testEmail).Return(nil)
	originalUserService := users.SetUsersService(mockUserService)

	data, _ := json.Marshal(loginForm)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(string(data)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp := httptest.NewRecorder()
	context := testServer.NewContext(req, resp)
	httpError := authController.Login(context).(*echo.HTTPError)
	assert.Equal(t, http.StatusUnauthorized, httpError.Code)
	users.SetUsersService(originalUserService)
}

func TestLoginFailWithInvalidPassword(t *testing.T) {
	testServer := echo.New()
	testServer.Validator = &common.CustomValidator{Validator: validator.New()}
	authController := AuthController{}
	var loginForm LoginRequest
	loginForm.Email = testEmail
	loginForm.Password = "wrong password"
	mockUserService := &MocksUsers.UsersService{}
	mockUserService.On("FindUserByEmail", testEmail).Return(&UserModels.User{
		Name:     testEmail,
		Password: testPassword,
		Role:     common.Admin,
	})
	originalUserService := users.SetUsersService(mockUserService)

	data, _ := json.Marshal(loginForm)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(string(data)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp := httptest.NewRecorder()
	context := testServer.NewContext(req, resp)

	httpError := authController.Login(context).(*echo.HTTPError)
	assert.Equal(t, http.StatusUnauthorized, httpError.Code)
	users.SetUsersService(originalUserService)
}

func TestLoginSuccess(t *testing.T) {
	test.InitTest()

	// create a test user
	user := UserModels.User{
		Name:     testName,
		Email:    testEmail,
		Role:     common.Admin,
		Password: testPassword,
	}
	mockUserService := &MocksUsers.UsersService{}
	mockUserService.On("FindUserByEmail", testEmail).Return(&user)
	originalUserService := users.SetUsersService(mockUserService)

	mockPasswordUtil := MocksUtils.PasswordUtil{}
	mockPasswordUtil.On("CheckPasswordHash", testPassword, testPassword).Return(true)
	originalPasswordUtil := utils.SetPasswordUtil(&mockPasswordUtil)

	testServer := echo.New()
	testServer.Validator = &common.CustomValidator{Validator: validator.New()}
	authController := AuthController{}
	var loginForm LoginRequest
	loginForm.Email = testEmail
	loginForm.Password = testPassword
	data, _ := json.Marshal(loginForm)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(string(data)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp := httptest.NewRecorder()
	context := testServer.NewContext(req, resp)

	if assert.NoError(t, authController.Login(context)) {
		assert.Equal(t, http.StatusOK, resp.Code)
	}
	users.SetUsersService(originalUserService)
	utils.SetPasswordUtil(originalPasswordUtil)
}
