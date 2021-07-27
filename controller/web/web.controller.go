package web

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type WebController struct {
}

func (webController WebController) Init(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "list.html", nil)
	})
	e.GET("/list", func(c echo.Context) error {
		return c.Render(http.StatusOK, "list.html", nil)
	})
	e.GET("/detail/:id", func(c echo.Context) error {
		return c.Render(http.StatusOK, "detail.html", map[string]interface{}{"id": c.Param("id")})
	})
	e.GET("/card", func(c echo.Context) error {
		return c.Render(http.StatusOK, "card.html", nil)
	})
	e.GET("/user", func(c echo.Context) error {
		return c.Render(http.StatusOK, "user.html", nil)
	})

	e.GET("/delete", func(c echo.Context) error {
		return c.Render(http.StatusOK, "delete.html", nil)
	})
}
