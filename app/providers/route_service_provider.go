package providers

import (
	"go-framework/framework/app"
	"go-framework/routes"
)

type RouteServiceProvider struct{}

func (p *RouteServiceProvider) Register(a *app.App) error {
	return nil
}

func (p *RouteServiceProvider) Boot(a *app.App) error {
	routes.Web(a.Router)
	routes.API(a.Router)
	return nil
}
