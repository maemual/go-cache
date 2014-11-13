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

func (item *Item) Expired() bool {
	if item.Expiration == nil {
		return false
	}
	return item.Expiration.Before(time.Now())
}

func New() *Cache {
	c := &Cache{
		items: map[string]*Item{},
	}
	go func() {
		for {
			time.Sleep(1 * time.Second)
			for k, v := range c.items {
				if v.Expired() {
					c.Delete(k)
				}
			}
		}
	}()
	return c
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	item, ok := c.items[key]
	if !ok || item.Expired() {
		c.RUnlock()
		return nil, false
	}
	c.RUnlock()
	return item.Object, true
}

func (c *Cache) Set(key string, val interface{}, dur time.Duration) {
	var t *time.Time
	if dur > 0 {
		tmp := time.Now().Add(dur)
		t = &tmp
	}
	c.Lock()
	c.items[key] = &Item{
		Object:     val,
		Expiration: t,
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
