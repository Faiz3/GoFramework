package bootstrap

import (
	"go-framework/app/providers"
	"go-framework/framework/app"
	"go-framework/framework/config"
	"go-framework/framework/logging"
)

func NewApp() *app.App {
	cfg := config.New()

	logging.New()

	application := app.New(cfg)

	application.Register(&providers.AppServiceProvider{})
	application.Register(&providers.DatabaseServiceProvider{})
	application.Register(&providers.ViewServiceProvider{})
	application.Register(&providers.RouteServiceProvider{})

	if err := application.Bootstrap(); err != nil {
		logging.Fatal("Failed to bootstrap application: %v", err)
	}

	return application
}
