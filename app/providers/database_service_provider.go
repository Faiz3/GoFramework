package providers

import (
	"go-framework/framework/app"
	"go-framework/framework/database"
)

type DatabaseServiceProvider struct{}

func (p *DatabaseServiceProvider) Register(a *app.App) error {
	return nil
}

func (p *DatabaseServiceProvider) Boot(a *app.App) error {
	a.DB = database.New(a.Config)
	return nil
}
