package lru_int_pool

import (
	"fmt"
	"testing"
	"time"
)

func TestLenProtected(t *testing.T) {
	cache := NewCache(2)

	for i := 0; i < 10; i++ {
		cache.Add(i, i)
	}
	if cache.Len() != 2 {
		t.Fatalf("Expected length 2, got %d", cache.Len())
	}

	if _, ok := cache.Get(8); !ok {
		t.Fatalf("Expected key2 to be present")
	}
	if _, ok := cache.Get(9); !ok {
		t.Fatalf("Expected key3 to be present")
	}
}

func TestCacheEviction(t *testing.T) {
	cache := NewCache(2)

	cache.Add(1, 1)
	cache.Add(2, 2)
	if cache.Len() != 2 {
		t.Fatalf("Expected length 2, got %d", cache.Len())
	}

	cache.Add(3, 3)
	if cache.Len() != 2 {
		t.Fatalf("Expected length 2, got %d", cache.Len())
	}

	if _, ok := cache.Get(1); ok {
		t.Fatalf("Expected key1 to be evicted")
	}
	if cache.Len() != 2 {
		t.Fatalf("Expected length 2, got %d", cache.Len())
	}

	if _, ok := cache.Get(2); !ok {
		t.Fatalf("Expected key2 to be present")
	}
	if cache.Len() != 2 {
		t.Fatalf("Expected length 2, got %d", cache.Len())
	}

	if v, ok := cache.Get(2); !ok {
		t.Fatalf("Expected key2 to be present %v", v)
	}

	if cache.Len() != 2 {
		t.Fatalf("Expected length 2, got %d", cache.Len())
	}

	if _, ok := cache.Get(2); !ok {
		t.Fatalf("Expected key2 to be present")
	}
	if cache.Len() != 2 {
		t.Fatalf("Expected length 2, got %d", cache.Len())
	}

	if _, ok := cache.Get(3); !ok {
		t.Fatalf("Expected key3 to be present")
	}
	if cache.Len() != 2 {
		t.Fatalf("Expected length 2, got %d", cache.Len())
	}
}

func TestCache(t *testing.T) {
	c := NewCache(3)

	c.Add(1, 1)
	c.Add(2, 2)
	c.Add(3, 3)

	// printCacheState(c)

	if v, ok := c.Get(1); !ok || v != 1 {
		t.Fatalf("Expected value1, got %v", v)
	}
	if v, ok := c.Get(2); !ok || v != 2 {
		t.Fatalf("Expected value2, got %v", v)
	}
	if v, ok := c.Get(3); !ok || v != 3 {
		t.Fatalf("Expected value3, got %v", v)
	}

	c.Add(4, 4)

	if v, ok := c.Get(1); ok {
		t.Fatalf("Expected key1 to be evicted, but got %v", v)
	}

	if v, ok := c.Get(2); !ok || v != 2 {
		t.Fatalf("Expected value2, got %v", v)
	}
	if v, ok := c.Get(3); !ok || v != 3 {
		t.Fatalf("Expected value3, got %v", v)
	}
	if v, ok := c.Get(4); !ok || v != 4 {
		t.Fatalf("Expected value4, got %v", v)
	}

	c.Add(2, 123)
	if v, ok := c.Get(2); !ok || v != 123 {
		t.Fatalf("Expected new_value2, got %v", v)
	}

	c.AddWithTTL(5, 5, 1*time.Millisecond)
	time.Sleep(2 * time.Millisecond)

	if v, ok := c.Get(5); ok {
		t.Fatalf("Expected key5 to be evicted, but got %v", v)
	}
}

func TestCache2(t *testing.T) {
	cache := NewCache(2)

	cache.Add(1, 1)
	val, ok := cache.Get(1)
	if !ok || val != 1 {
		t.Fatalf("Expected value1, got %v", val)
	}

	cache.Add(2, 2)
	val, ok = cache.Get(2)
	if !ok || val != 2 {
		t.Fatalf("Expected value2, got %v", val)
	}

	cache.Add(3, 3)
	val, ok = cache.Get(1)
	if ok {
		t.Fatalf("Expected key1 to be evicted, got %v", val)
	}

	val, ok = cache.Get(2)
	if !ok || val != 2 {
		t.Fatalf("Expected value2, got %v", val)
	}

	val, ok = cache.Get(3)
	if !ok || val != 3 {
		t.Fatalf("Expected value3, got %v", val)
	}

	cache.AddWithTTL(4, 4, 1*time.Second)
	val, ok = cache.Get(4)
	if !ok || val != 4 {
		t.Fatalf("Expected value4, got %v", val)
	}

	time.Sleep(2 * time.Second)

	val, ok = cache.Get(4)
	if ok {
		t.Fatalf("Expected key4 to be evicted, got %v", val)
	}
}

func TestCacheWithTTLAndNoTTL(t *testing.T) {
	cache := NewCache(2)

	cache.AddWithTTL(1, 1, 2*time.Second)
	val, ok := cache.Get(1)
	if !ok || val != 1 {
		t.Fatalf("Expected value1, got %v", val)
	}

	time.Sleep(1 * time.Second)

	val, ok = cache.Get(1)
	if !ok || val != 1 {
		t.Fatalf("Expected value1, got %v", val)
	}

	time.Sleep(2 * time.Second)

	val, ok = cache.Get(1)
	if ok {
		t.Fatalf("Expected key1 to be evicted, got %v", val)
	}

	cache.Add(2, 2)
	val, ok = cache.Get(2)
	if !ok || val != 2 {
		t.Fatalf("Expected value2, got %v", val)
	}

	cache.Add(3, 3)
	val, ok = cache.Get(2)
	if !ok || val != 2 {
		t.Fatalf("Expected value2, got %v", val)
	}
}

func printCacheState(c *Cache) {
	fmt.Println("Cache items:")
	for k, v := range c.items {
		fmt.Printf("Key: %v, Value: %v, Expiration: %v\n", k, v.Value, v.Expiration)
	}
	fmt.Println("Cache list order:")
	node := c.list.Head
	for node != nil {
		fmt.Printf("Key: %v, Value: %v\n", node.Key, node.Value)
		node = node.Next
	}
	fmt.Println("--------")
}
