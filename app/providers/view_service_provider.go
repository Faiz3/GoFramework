package providers

import "go-framework/framework/app"

type ViewServiceProvider struct{}

func (p *ViewServiceProvider) Register(a *app.App) error {
	return nil
}

func (p *ViewServiceProvider) Boot(a *app.App) error {
	return nil
}
