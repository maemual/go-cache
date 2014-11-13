package cache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := New()
	_, found := c.Get("xx")
	if found {
		t.Error("WTF?")
	}

	c.Set("key", "lalala", -1)
	val, found := c.Get("key")
	if !found {
		t.Error("WHY???")
	}
	if val != "lalala" {
		t.Error("FUCK!")
	}
	c.Delete("key")
	_, found = c.Get("key")
	if found {
		t.Error("impossiable!")
	}
	c.Set("key", "bababa", -1)
	c.Flush()
	_, found = c.Get("key")
	if found {
		t.Error("impossiable again!")
	}

	c.Set("key", "val", 2*time.Second)
	_, found = c.Get("key")
	if !found {
		t.Error("must have this")
	}
	time.Sleep(2 * time.Second)
	_, found = c.Get("key")
	if found {
		t.Error("must not have this")
	}
}
