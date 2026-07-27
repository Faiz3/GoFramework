package routing

import (
	"go-framework/framework/http"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	fiber *fiber.App
	group fiber.Router
}

func NewRouter(f *fiber.App) *Router {
	return &Router{
		fiber: f,
		group: f,
	}
}

func (r *Router) Group(prefix string, callback func(router *Router), middleware ...fiber.Handler) {
	g := r.group.Group(prefix, middleware...)
	sub := &Router{fiber: r.fiber, group: g}
	callback(sub)
}

func (r *Router) Get(path string, handler interface{}, middleware ...fiber.Handler) {
	r.addRoute("GET", path, handler, middleware...)
}

func (r *Router) Post(path string, handler interface{}, middleware ...fiber.Handler) {
	r.addRoute("POST", path, handler, middleware...)
}

func (r *Router) Put(path string, handler interface{}, middleware ...fiber.Handler) {
	r.addRoute("PUT", path, handler, middleware...)
}

func (r *Router) Delete(path string, handler interface{}, middleware ...fiber.Handler) {
	r.addRoute("DELETE", path, handler, middleware...)
}

func (r *Router) Patch(path string, handler interface{}, middleware ...fiber.Handler) {
	r.addRoute("PATCH", path, handler, middleware...)
}

func (r *Router) Options(path string, handler interface{}, middleware ...fiber.Handler) {
	r.addRoute("OPTIONS", path, handler, middleware...)
}

func (r *Router) Resource(name string, controller http.ResourceController, middleware ...fiber.Handler) {
	r.Group("/"+name, func(router *Router) {
		router.Get("/", controller.Index, middleware...)
		router.Post("/", controller.Store, middleware...)
		router.Get("/:id", controller.Show, middleware...)
		router.Put("/:id", controller.Update, middleware...)
		router.Delete("/:id", controller.Destroy, middleware...)
	}, middleware...)
}

func (r *Router) Static(prefix, root string) {
	r.fiber.Static(prefix, root)
}

func (r *Router) addRoute(method, path string, handler interface{}, middleware ...fiber.Handler) {
	var h fiber.Handler

	switch hdl := handler.(type) {
	case func(*fiber.Ctx) error:
		h = hdl
	case http.Controller:
		h = hdl.Handle
	case func() error:
		h = func(c *fiber.Ctx) error { return hdl() }
	default:
		h = func(c *fiber.Ctx) error { return c.SendString("handler not found") }
	}

	handlers := make([]fiber.Handler, 0)
	handlers = append(handlers, middleware...)
	handlers = append(handlers, h)

	switch method {
	case "GET":
		r.group.Get(path, handlers...)
	case "POST":
		r.group.Post(path, handlers...)
	case "PUT":
		r.group.Put(path, handlers...)
	case "DELETE":
		r.group.Delete(path, handlers...)
	case "PATCH":
		r.group.Patch(path, handlers...)
	case "OPTIONS":
		r.group.Options(path, handlers...)
	}
}
