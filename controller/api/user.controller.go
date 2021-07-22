package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	_ "platform-sample/docs"
	"platform-sample/model"
	"platform-sample/service"
	"strconv"
)

type UserController struct {
	service.UserService
}

func (UserController) NewUserController(service service.UserService) *UserController {
	return &UserController{service}
}

func (userController *UserController) Init(e *echo.Group) {
	e.POST("", userController.CreateUser)
	e.GET("", userController.GetUsers)
	e.DELETE("/:id", userController.DeleteUser)
	e.GET("/:id", userController.GetUser)
	e.PATCH("", userController.UpdateUser)
}

// CreateUser is aaaa
// @Summary Create user
// @Description Create new user
// @Accept json
// @Produce json
// @Param user body UserDto true "body of the user"
// @Success 201 {object} model.User
// @Router /users [post]
func (userController *UserController) CreateUser(c echo.Context) error {
	userDto := &UserDto{}
	bindErr := c.Bind(userDto)
	user := userDto.toModel()

	if bindErr != nil {
		c.Logger().Error(bindErr)
		return ReturnApiFail(c, http.StatusBadRequest, ApiParameterError, bindErr)
	}
	createUser, err := userController.UserService.CreateUser(user)
	if err != nil {
		c.Logger().Error(err)
		return ReturnApiFail(c, http.StatusInternalServerError, ApiQueryError, err)
	}
	c.Logger().Info(createUser)
	return ReturnApiSuccess(c, http.StatusCreated, createUser)
}

// GetUsers is aaa
// @Summary Get all users
// @Description Get all user's info
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Router /users [get]
func (userController *UserController) GetUsers(c echo.Context) error {
	users, err := userController.UserService.GetUsers()
	if err != nil {
		c.Logger().Error(err)
		return ReturnApiFail(c, http.StatusInternalServerError, ApiQueryError, err)
	}
	c.Logger().Info(users)
	return ReturnApiSuccess(c, http.StatusOK, users)
}

// DeleteUser aaa
// @Summary Delete user
// @Description Delete user's info
// @Accept json
// @Produce json
// @Param id path string true "id of the user"
// @Success 204 {object} model.User
// @Router /users/{id} [delete]
func (userController *UserController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return ReturnApiFail(c, http.StatusBadRequest, ApiParameterError, err)
	}

	err = userController.UserService.DeleteUser(id)
	if err != nil {
		c.Logger().Error(err)
		return ReturnApiFail(c, http.StatusInternalServerError, ApiQueryError, err)
	}
	return ReturnApiSuccess(c, http.StatusNoContent, nil)
}

// GetUser aaa
// @Summary Get user
// @Description Get user's info
// @Accept json
// @Produce json
// @Param id path string true "id of the user"
// @Success 200 {object} model.User
// @Router /users/{id} [get]
func (userController *UserController) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return ReturnApiFail(c, http.StatusBadRequest, ApiParameterError, err)
	}

	user, err := userController.UserService.GetUser(id)
	if err != nil {
		c.Logger().Error(err)
		return ReturnApiFail(c, http.StatusInternalServerError, ApiQueryError, err)
	}
	c.Logger().Info(user)
	return ReturnApiSuccess(c, http.StatusOK, user)
}

// UpdateUser aaa
// @Summary Update user
// @Description Get user's info
// @Accept json
// @Produce json
// @Param name body model.User true "body of the user"
// @Success 201 {object} model.User
// @Router /users [patch]
func (userController *UserController) UpdateUser(c echo.Context) error {
	user := &model.User{}
	bindErr := c.Bind(user)
	if bindErr != nil {
		c.Logger().Error(bindErr)
		return ReturnApiFail(c, http.StatusBadRequest, ApiParameterError, bindErr)
	}

	createUser, err := userController.UserService.UpdateUser(user)
	if err != nil {
		c.Logger().Error(bindErr)
		return ReturnApiFail(c, http.StatusInternalServerError, ApiQueryError, err)
	}
	c.Logger().Info(createUser)
	return ReturnApiSuccess(c, http.StatusOK, user)
}
