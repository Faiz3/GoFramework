package app

import (
	"go-framework/framework/config"
	"go-framework/framework/database"
	"go-framework/framework/routing"
	"go-framework/framework/view"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type ServiceProvider interface {
	Register(app *App) error
	Boot(app *App) error
}

type App struct {
	Fiber     *fiber.App
	Config    *config.Config
	Router    *routing.Router
	DB        *database.Database
	View      *view.View
	providers []ServiceProvider
}

func New(cfg *config.Config) *App {
	viewEngine := html.New("resources/views", ".html")

	f := fiber.New(fiber.Config{
		AppName:      cfg.GetString("app.name"),
		ReadTimeout:  cfg.GetDuration("app.read_timeout"),
		WriteTimeout: cfg.GetDuration("app.write_timeout"),
		BodyLimit:    cfg.GetInt("app.body_limit") * 1024 * 1024,
		Views:        viewEngine,
		ViewsLayout:  "",
	})

	return &App{
		Fiber:     f,
		Config:    cfg,
		Router:    routing.NewRouter(f),
		View:      &view.View{Engine: viewEngine},
		providers: make([]ServiceProvider, 0),
	}
}

func (a *App) Register(provider ServiceProvider) {
	a.providers = append(a.providers, provider)
}

func (a *App) Bootstrap() error {
	for _, p := range a.providers {
		if err := p.Register(a); err != nil {
			return err
		}
	}
	for _, p := range a.providers {
		if err := p.Boot(a); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) Run(addr string) error {
	return a.Fiber.Listen(addr)
}
