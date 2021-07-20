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
	e.POST("", userController.createUser)
	e.GET("", userController.getUsers)
	e.DELETE("/:id", userController.deleteUser)
	e.GET("/:id", userController.getUser)
	e.PATCH("", userController.updateUser)
}

// @Summary Create user
// @Description Create new user
// @Accept json
// @Produce json
// @Param user body model.User true "body of the user"
// @Success 203 {object} model.User
// @Router /users [post]
func (userController *UserController) createUser(c echo.Context) error {
	user := &model.User{}
	bindErr := c.Bind(user)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, bindErr)
	}

	createUser, err := userController.UserService.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, createUser)
}

// @Summary Get all users
// @Description Get all user's info
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Router /users [get]
func (userController *UserController) getUsers(c echo.Context) error {
	users, err := userController.UserService.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

// @Summary Delete user
// @Description Delete user's info
// @Accept json
// @Produce json
// @Param id path string true "id of the user"
// @Success 204 {object} model.User
// @Router /users/{id} [delete]
func (userController *UserController) deleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = userController.UserService.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK,nil)
}

// @Summary Get user
// @Description Get user's info
// @Accept json
// @Produce json
// @Param id path string true "id of the user"
// @Success 200 {object} model.User
// @Router /users/{id} [get]
func (userController *UserController) getUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := userController.UserService.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK,user)
}

// @Summary Update user
// @Description Get user's info
// @Accept json
// @Produce json
// @Param name body model.User true "body of the user"
// @Success 201 {object} model.User
// @Router /users [patch]
func (userController *UserController) updateUser(c echo.Context) error {
	user := &model.User{}
	bindErr := c.Bind(user)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, bindErr)
	}

	createUser, err := userController.UserService.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, createUser)
}
