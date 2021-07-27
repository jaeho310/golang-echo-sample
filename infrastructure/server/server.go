package server

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"html/template"
	"io"
	"platform-sample/controller/api"
	"platform-sample/controller/web"
	_ "platform-sample/docs"
	"platform-sample/repository"
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
	cardController := server.InjectCardController()
	cardController.Init(e.Group("/api/cards"))

	userController := server.InjectUserController()
	userController.Init(e.Group("/api/users"))

	// swagger setting
	e.GET("/swagger/*", echoSwagger.WrapHandler)

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

func (server Server) InjectDb() *gorm.DB {
	return server.MainDb
}

func (server Server) InjectUserRepository() (*repository.UserRepositoryImpl, *repository.CardRepositoryImpl) {
	return repository.UserRepositoryImpl{}.NewUserRepositoryImpl(server.InjectDb()),
		repository.CardRepositoryImpl{}.NewCardRepositoryImpl(server.InjectDb())
	// TODO 이거 같은객체가 아니다. 동시성 및 멤버에 문제가 생길가능성 존재하여 위험하다.
	// 방법 1. db 쿼리문을 재사용하지않고 매번 사용한다. -> 위험하고 db단의 트랜잭션 처리를 애플리케이션단에서 잡아줘야한다.
	// 방법 2. 의존성 주입을 해줄 객체를 sington이나 static 객체로 한곳에서 관리한다. -> 주입부분을 갈아끼워야한다.
	// echo를 깊게 공부하지않아서 정확하지는 않지만, 하나의 transactin은 같은객체라면 echo가 잡아놔서 동시성을 처리해줄텐데
	// 어떤방법이 좋을지 모르겠다.
}

func (server Server) InjectUserService() *service.UserServiceImpl {
	return service.UserServiceImpl{}.NewUserServiceImpl(server.InjectUserRepository())
}

func (server Server) InjectUserController() *api.UserController {
	return api.UserController{}.NewUserController(server.InjectUserService())
}

func (server Server) InjectCardRepository() *repository.CardRepositoryImpl {
	return repository.CardRepositoryImpl{}.NewCardRepositoryImpl(server.InjectDb())
}

func (server Server) InjectCardService() *service.CardServiceImpl {
	return service.CardServiceImpl{}.NewCardServiceImpl(server.InjectCardRepository())
}

func (server Server) InjectCardController() *api.CardController {
	return api.CardController{}.NewCardController(server.InjectCardService())
}
