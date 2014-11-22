package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	items             map[string]*Item
	defaultExpiration time.Duration
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

func New(defaultExpiration, cleanInterval time.Duration) *Cache {
	c := &Cache{
		items:             map[string]*Item{},
		defaultExpiration: defaultExpiration,
	}
	if cleanInterval > 0 {
		go func() {
			for {
				time.Sleep(cleanInterval)
				for k, v := range c.items {
					if v.Expired() {
						c.Delete(k)
					}
				}
			}
		}()
	}
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
	c.Lock()
	if dur == 0 {
		dur = c.defaultExpiration
	}
	if dur > 0 {
		tmp := time.Now().Add(dur)
		t = &tmp
	}
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

func (c *Cache) Increment(key string, x int64) error {
	c.Lock()
	val, ok := c.items[key]
	if !ok || val.Expired() {
		c.Unlock()
		return fmt.Errorf("Item %s not found", key)
	}
	switch val.Object.(type) {
	case int:
		val.Object = val.Object.(int) + int(x)
	case int8:
		val.Object = val.Object.(int8) + int8(x)
	case int16:
		val.Object = val.Object.(int16) + int16(x)
	case int32:
		val.Object = val.Object.(int32) + int32(x)
	case int64:
		val.Object = val.Object.(int64) + x
	case uint:
		val.Object = val.Object.(uint) + uint(x)
	case uint8:
		val.Object = val.Object.(uint8) + uint8(x)
	case uint16:
		val.Object = val.Object.(uint16) + uint16(x)
	case uint32:
		val.Object = val.Object.(uint32) + uint32(x)
	case uint64:
		val.Object = val.Object.(uint64) + uint64(x)
	case uintptr:
		val.Object = val.Object.(uintptr) + uintptr(x)
	default:
		c.Unlock()
		return fmt.Errorf("The value type error")
	}
	c.Unlock()
	return nil
}

func (c *Cache) Decrement(key string, x int64) error {
	c.Lock()
	val, ok := c.items[key]
	if !ok || val.Expired() {
		c.Unlock()
		return fmt.Errorf("Item %s not found", key)
	}
	switch val.Object.(type) {
	case int:
		val.Object = val.Object.(int) - int(x)
	case int8:
		val.Object = val.Object.(int8) - int8(x)
	case int16:
		val.Object = val.Object.(int16) - int16(x)
	case int32:
		val.Object = val.Object.(int32) - int32(x)
	case int64:
		val.Object = val.Object.(int64) - x
	case uint:
		val.Object = val.Object.(uint) - uint(x)
	case uint8:
		val.Object = val.Object.(uint8) - uint8(x)
	case uint16:
		val.Object = val.Object.(uint16) - uint16(x)
	case uint32:
		val.Object = val.Object.(uint32) - uint32(x)
	case uint64:
		val.Object = val.Object.(uint64) - uint64(x)
	case uintptr:
		val.Object = val.Object.(uintptr) - uintptr(x)
	default:
		c.Unlock()
		return fmt.Errorf("The value type error")
	}
	c.Unlock()
	return nil
}

func (c *Cache) ItemCount() int {
	c.RLock()
	counts := len(c.items)
	c.RUnlock()
	return counts
}
