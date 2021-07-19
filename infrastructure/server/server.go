package server

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"platform-sample/controller/api"
	"platform-sample/controller/web"
	"platform-sample/infrastructure/database"
	"platform-sample/service"
)

type Server struct {
	MainDb *gorm.DB
}

type TemplateRenderer struct {
	templates *template.Template
}

func (server Server) Init() {
	e := echo.New()

	// web controller setting
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("view/templates/*.html")),
	}
	e.Renderer = renderer
	e.Static("/static", "view/static")
	web.WebController{}.Init(e)

	// api controller setting
	userController := server.injectUserController()
	userController.Init(e.Group("/api/users"))

	e.Logger.Fatal(e.Start(":8395"))
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func (server Server) contextDB(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}

func (server Server) injectDb() *gorm.DB {
	return server.MainDb
}

func (server Server) injectUserRepository() *database.UserRepositoryImpl {
	return database.UserRepositoryImpl{}.NewUserRepositoryImpl(server.injectDb())
}

func (server Server) injectUserService() *service.UserServiceImpl {
	return service.UserServiceImpl{}.NewUserServiceImpl(server.injectUserRepository())
}

func (server Server) injectUserController() *api.UserController {
	return api.UserController{}.NewUserController(server.injectUserService())
}

