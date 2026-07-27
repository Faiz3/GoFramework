package routes

import (
	"go-framework/app/http/controllers"
	"go-framework/app/http/middleware"
	"go-framework/framework/routing"
)

func API(router *routing.Router) {
	router.Group("/api", func(r *routing.Router) {
		r.Post("/auth/login", &controllers.AuthController{})
		r.Post("/auth/register", &controllers.AuthController{})

		r.Group("/protected", func(pr *routing.Router) {
			pr.Get("/profile", &controllers.UserController{})
		}, middleware.JWTAuth)
	})
}
