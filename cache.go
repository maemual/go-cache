package cache

import (
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	items map[string]*Item
}

type Item struct {
	Object     interface{}
	Expiration *time.Time
}

func New() *Cache {
	return &Cache{
		items: map[string]*Item{},
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	item, ok := c.items[key]
	if !ok {
		c.RUnlock()
		return nil, false
	}
	c.RUnlock()
	return item.Object, true
}

func (c *Cache) Set(key string, val interface{}) {
	c.Lock()
	c.items[key] = &Item{
		Object: val,
	}
	c.Unlock()
}

func (c *Cache) Delete(key string) {
	c.Lock()
	delete(c.items, key)
	c.Unlock()
}

func (c *Cache) Flush() {
	c.Lock()
	c.items = map[string]*Item{}
	c.Unlock()
}
