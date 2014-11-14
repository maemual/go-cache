package cache

import (
	"sync"
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
	cnt := c.ItemCount()
	if cnt != 1 {
		t.Error("The number of cache must be 1")
	}
	c.Flush()
	_, found = c.Get("key")
	if found {
		t.Error("All keys are flush, you should not get")
	}
	cnt = c.ItemCount()
	if cnt != 0 {
		t.Error("The number of cache must be 0")
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

func BenchmarkCacheGet(b *testing.B) {
	b.StopTimer()
	tc := New()
	tc.Set("key", "values", -1)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Get("key")
	}
}

func BenchmarkRWMutexGet(b *testing.B) {
	b.StopTimer()
	m := map[string]string{
		"key": "values",
	}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.RLock()
		_, _ = m["key"]
		mu.RUnlock()
	}
}

func BenchmarkCacheSet(b *testing.B) {
	b.StopTimer()
	tc := New()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Set("key", "values", -1)
	}
}

func BenchmarkRWMutexSet(b *testing.B) {
	b.StopTimer()
	m := map[string]string{}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		m["key"] = "values"
		mu.Unlock()
	}
}

func BenchmarkCacheIncrement(b *testing.B) {
	b.StopTimer()
	tc := New()
	tc.Set("key", 0, -1)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Increment("key", 1)
	}
}

func BenchmarkCacheDecrement(b *testing.B) {
	b.StopTimer()
	tc := New()
	tc.Set("key", 0, -1)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Decrement("key", 1)
	}
}
