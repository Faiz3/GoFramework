package providers

import (
	"go-framework/framework/app"
	"go-framework/framework/auth"
)

type AuthServiceProvider struct{}

func (p *AuthServiceProvider) Register(a *app.App) error {
	return nil
}

func (p *AuthServiceProvider) Boot(a *app.App) error {
	secret := a.Config.Get("JWT_SECRET", "default-secret")
	_ = auth.New(secret)
	return nil
}
