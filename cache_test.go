package cache

import (
	"testing"
)

func TestCache(t *testing.T) {
	c := New()
	_, found := c.Get("xx")
	if found {
		t.Error("WTF?")
	}

	c.Set("key", "lalala")
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
	c.Set("key", "bababa")
	c.Flush()
	_, found = c.Get("key")
	if found {
		t.Error("impossiable again!")
	}
}
