package view

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type View struct {
	Engine *html.Engine
}

func New(viewsDir string, extension string) *View {
	engine := html.New(viewsDir, extension)
	engine.Reload(true)
	return &View{Engine: engine}
}

func (v *View) RegisterGlobal(key string, value interface{}) {
	v.Engine.AddFunc(key, func() interface{} {
		return value
	})
}

func (v *View) Share(key string, value interface{}) {
	v.RegisterGlobal(key, value)
}

func (v *View) Exists(name string) bool {
	return v.Engine != nil
}

func (v *View) Render(c *fiber.Ctx, name string, data fiber.Map) error {
	return c.Render(name, data)
}
