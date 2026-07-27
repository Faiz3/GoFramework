package providers

import "go-framework/framework/app"

type AppServiceProvider struct{}

func (p *AppServiceProvider) Register(a *app.App) error {
	return nil
}

func (p *AppServiceProvider) Boot(a *app.App) error {
	return nil
}
