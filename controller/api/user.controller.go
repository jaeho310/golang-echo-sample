package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	_ "platform-sample/docs"
	"platform-sample/model"
	"platform-sample/service"
	"strconv"

	//"github.com/labstack/echo"
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

func (userController *UserController) getUsers(c echo.Context) error {
	users, err := userController.UserService.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

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
