package http

import (
	"github.com/gofiber/fiber/v2"
)

type Request struct {
	ctx *fiber.Ctx
}

func NewRequest(ctx *fiber.Ctx) *Request {
	return &Request{ctx: ctx}
}

func (r *Request) Input(key string, defaultValue ...string) string {
	val := r.ctx.Params(key)
	if val == "" {
		val = r.ctx.Query(key)
	}
	if val == "" {
		val = r.ctx.FormValue(key)
	}
	if val == "" {
		val = r.ctx.GetRespHeader(key)
	}
	if val == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return val
}

func (r *Request) All() map[string]string {
	params := make(map[string]string)
	for _, p := range r.ctx.Route().Params {
		params[p] = r.ctx.Params(p)
	}
	r.ctx.Request().URI().QueryArgs().VisitAll(func(key, value []byte) {
		params[string(key)] = string(value)
	})
	return params
}

func (r *Request) Only(keys ...string) map[string]string {
	all := r.All()
	result := make(map[string]string)
	for _, k := range keys {
		if v, ok := all[k]; ok {
			result[k] = v
		}
	}
	return result
}

func (r *Request) Has(key string) bool {
	return r.Input(key) != ""
}

func (r *Request) BearerToken() string {
	auth := r.ctx.Get("Authorization")
	if len(auth) > 7 && auth[:7] == "Bearer " {
		return auth[7:]
	}
	return ""
}

func (r *Request) IP() string {
	return r.ctx.IP()
}

func (r *Request) Method() string {
	return r.ctx.Method()
}

func (r *Request) Path() string {
	return r.ctx.Path()
}

func (r *Request) Validate(validation map[string]string) error {
	for field, rule := range validation {
		val := r.Input(field)
		if err := validateField(field, val, rule); err != nil {
			return err
		}
	}
	return nil
}

func validateField(field, value, rule string) error {
	return nil
}

type Validator interface {
	Rules() map[string]string
	Messages() map[string]string
	Validate(c *fiber.Ctx) error
}

type BaseRequest struct {
	ctx *fiber.Ctx
}

func (r *BaseRequest) Validate(c *fiber.Ctx) error {
	r.ctx = c
	return nil
}
