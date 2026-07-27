package session

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Session struct {
	store *session.Store
}

func New(key string) *Session {
	store := session.New(session.Config{
		KeyLookup: "cookie:" + key,
	})
	return &Session{store: store}
}

func (s *Session) Get(c *fiber.Ctx, key string) (interface{}, bool) {
	sess, err := s.store.Get(c)
	if err != nil {
		return nil, false
	}
	val := sess.Get(key)
	return val, val != nil
}

func (s *Session) Set(c *fiber.Ctx, key string, value interface{}) error {
	sess, err := s.store.Get(c)
	if err != nil {
		return err
	}
	sess.Set(key, value)
	return sess.Save()
}

func (s *Session) Forget(c *fiber.Ctx, key string) error {
	sess, err := s.store.Get(c)
	if err != nil {
		return err
	}
	sess.Delete(key)
	return sess.Save()
}

func (s *Session) Destroy(c *fiber.Ctx) error {
	sess, err := s.store.Get(c)
	if err != nil {
		return err
	}
	if err := sess.Destroy(); err != nil {
		return err
	}
	return sess.Save()
}

func (s *Session) Flash(c *fiber.Ctx, key string, value interface{}) error {
	return s.Set(c, key, value)
}

func (s *Session) Store() *session.Store {
	return s.store
}
