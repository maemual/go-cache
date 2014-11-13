package cache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := New()
	_, found := c.Get("xx")
	if found {
		t.Error("You should not get")
	}

	c.Set("key", "lalala", -1)
	val, found := c.Get("key")
	if !found {
		t.Error("You must get this value")
	}
	if val != "lalala" {
		t.Error("You get a wrong value")
	}
	c.Delete("key")
	_, found = c.Get("key")
	if found {
		t.Error("The key is delete, you should not get")
	}
	c.Set("key", "bababa", -1)
	c.Flush()
	_, found = c.Get("key")
	if found {
		t.Error("All keys are flush, you should not get")
	}

	c.Set("key", "val", 2*time.Second)
	_, found = c.Get("key")
	if !found {
		t.Error("must have this")
	}
	time.Sleep(2 * time.Second)
	_, found = c.Get("key")
	if found {
		t.Error("The key is time out, you should not get")
	}
}

func TestIncDec(t *testing.T) {
	c := New()
	c.Set("key", 100, -1)
	c.Increment("key", 1)
	val, _ := c.Get("key")
	if val != 101 {
		t.Error("Increment error")
	}
	c.Decrement("key", 1)
	val, _ = c.Get("key")
	if val != 100 {
		t.Error("Decrement error")
	}
}
