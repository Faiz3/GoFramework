package cache

import (
	"sync"
	"time"
)

type item struct {
	value     interface{}
	expiresAt time.Time
}

type Cache struct {
	mu    sync.RWMutex
	items map[string]*item
}

var instance *Cache

func New() *Cache {
	c := &Cache{
		items: make(map[string]*item),
	}
	instance = c
	go c.gc()
	return c
}

func GetInstance() *Cache {
	if instance == nil {
		instance = New()
	}
	return instance
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	it, ok := c.items[key]
	if !ok {
		return nil, false
	}

	if !it.expiresAt.IsZero() && time.Now().After(it.expiresAt) {
		delete(c.items, key)
		return nil, false
	}

	return it.value, true
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	expiresAt := time.Time{}
	if ttl > 0 {
		expiresAt = time.Now().Add(ttl)
	}

	c.items[key] = &item{
		value:     value,
		expiresAt: expiresAt,
	}
}

func (c *Cache) Forget(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

func (c *Cache) Has(key string) bool {
	_, ok := c.Get(key)
	return ok
}

func (c *Cache) Remember(key string, ttl time.Duration, callback func() (interface{}, error)) (interface{}, error) {
	if val, ok := c.Get(key); ok {
		return val, nil
	}

	val, err := callback()
	if err != nil {
		return nil, err
	}

	c.Set(key, val, ttl)
	return val, nil
}

func (c *Cache) Flush() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[string]*item)
}

func (c *Cache) gc() {
	ticker := time.NewTicker(5 * time.Minute)
	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for k, v := range c.items {
			if !v.expiresAt.IsZero() && now.After(v.expiresAt) {
				delete(c.items, k)
			}
		}
		c.mu.Unlock()
	}
}

func Get(key string) (interface{}, bool) {
	return GetInstance().Get(key)
}

func Set(key string, value interface{}, ttl time.Duration) {
	GetInstance().Set(key, value, ttl)
}

func Forget(key string) {
	GetInstance().Forget(key)
}
