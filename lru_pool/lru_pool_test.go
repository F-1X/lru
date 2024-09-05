package lru_pool

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

	cache.Add("key1", "value1")
	cache.Add("key2", "value2")
	if cache.Len() != 2 {
		t.Fatalf("Expected length 2, got %d", cache.Len())
	}

	cache.Add("key3", "value3")
	if cache.Len() != 2 {
		t.Fatalf("Expected length 2, got %d", cache.Len())
	}

	if _, ok := cache.Get("key1"); ok {
		t.Fatalf("Expected key1 to be evicted")
	}
	if cache.Len() != 2 {
		t.Fatalf("Expected length 2, got %d", cache.Len())
	}

	if _, ok := cache.Get("key2"); !ok {
		t.Fatalf("Expected key2 to be present")
	}
	if cache.Len() != 2 {
		t.Fatalf("Expected length 2, got %d", cache.Len())
	}
	
	if v, ok := cache.Get("key2"); !ok {
		t.Fatalf("Expected key2 to be present %v",v)
	}

	if cache.Len() != 2 {
		t.Fatalf("Expected length 2, got %d", cache.Len())
	}

	
	if _, ok := cache.Get("key2"); !ok {
		t.Fatalf("Expected key2 to be present")
	}
	if cache.Len() != 2 {
		t.Fatalf("Expected length 2, got %d", cache.Len())
	}




	if _, ok := cache.Get("key3"); !ok {
		t.Fatalf("Expected key3 to be present")
	}
	if cache.Len() != 2 {
		t.Fatalf("Expected length 2, got %d", cache.Len())
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
func TestCache(t *testing.T) {
	c := NewCache(3)

	c.Add("key1", "value1")
	c.Add("key2", "value2")
	c.Add("key3", "value3")

	printCacheState(c)

	// Проверяем, что все элементы добавлены
	if v, ok := c.Get("key1"); !ok || v != "value1" {
		t.Fatalf("Expected value1, got %v", v)
	}
	if v, ok := c.Get("key2"); !ok || v != "value2" {
		t.Fatalf("Expected value2, got %v", v)
	}
	if v, ok := c.Get("key3"); !ok || v != "value3" {
		t.Fatalf("Expected value3, got %v", v)
	}

	c.Add("key4", "value4")

	if v, ok := c.Get("key1"); ok {
		t.Fatalf("Expected key1 to be evicted, but got %v", v)
	}

	if v, ok := c.Get("key2"); !ok || v != "value2" {
		t.Fatalf("Expected value2, got %v", v)
	}
	if v, ok := c.Get("key3"); !ok || v != "value3" {
		t.Fatalf("Expected value3, got %v", v)
	}
	if v, ok := c.Get("key4"); !ok || v != "value4" {
		t.Fatalf("Expected value4, got %v", v)
	}

	c.Add("key2", "new_value2")
	if v, ok := c.Get("key2"); !ok || v != "new_value2" {
		t.Fatalf("Expected new_value2, got %v", v)
	}

	c.AddWithTTL("key5", "value5", 1*time.Millisecond)
	time.Sleep(2 * time.Millisecond)

	if v, ok := c.Get("key5"); ok {
		t.Fatalf("Expected key5 to be evicted, but got %v", v)
	}
}

func TestCache2(t *testing.T) {
	cache := NewCache(2)

	cache.Add("key1", "value1")
	val, ok := cache.Get("key1")
	if !ok || val != "value1" {
		t.Fatalf("Expected value1, got %v", val)
	}

	cache.Add("key2", "value2")
	val, ok = cache.Get("key2")
	if !ok || val != "value2" {
		t.Fatalf("Expected value2, got %v", val)
	}

	cache.Add("key3", "value3")
	val, ok = cache.Get("key1")
	if ok {
		t.Fatalf("Expected key1 to be evicted, got %v", val)
	}

	val, ok = cache.Get("key2")
	if !ok || val != "value2" {
		t.Fatalf("Expected value2, got %v", val)
	}

	val, ok = cache.Get("key3")
	if !ok || val != "value3" {
		t.Fatalf("Expected value3, got %v", val)
	}

	cache.AddWithTTL("key4", "value4", 1*time.Second)
	val, ok = cache.Get("key4")
	if !ok || val != "value4" {
		t.Fatalf("Expected value4, got %v", val)
	}

	time.Sleep(2 * time.Second)

	val, ok = cache.Get("key4")
	if ok {
		t.Fatalf("Expected key4 to be evicted, got %v", val)
	}
}

func TestCacheWithTTLAndNoTTL(t *testing.T) {
	cache := NewCache(2)

	cache.AddWithTTL("key1", "value1", 2*time.Second)
	val, ok := cache.Get("key1")
	if !ok || val != "value1" {
		t.Fatalf("Expected value1, got %v", val)
	}

	time.Sleep(1 * time.Second)

	val, ok = cache.Get("key1")
	if !ok || val != "value1" {
		t.Fatalf("Expected value1, got %v", val)
	}

	time.Sleep(2 * time.Second)

	val, ok = cache.Get("key1")
	if ok {
		t.Fatalf("Expected key1 to be evicted, got %v", val)
	}

	cache.Add("key2", "value2")
	val, ok = cache.Get("key2")
	if !ok || val != "value2" {
		t.Fatalf("Expected value2, got %v", val)
	}

	cache.Add("key3", "value3")
	val, ok = cache.Get("key2")
	if !ok || val != "value2" {
		t.Fatalf("Expected value2, got %v", val)
	}
}