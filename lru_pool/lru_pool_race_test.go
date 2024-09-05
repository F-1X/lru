package lru_pool

import (
	"sync"
	"testing"
	"time"
)

func TestCacheRaceConditionParallel(t *testing.T) {
	cache := NewCache(100)

	t.Run("Add and Get parallel", func(t *testing.T) {
		t.Parallel()
		var wg sync.WaitGroup

		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				key := i
				value := i * 10

				cache.AddWithTTL(key, value, time.Minute)

				val, ok := cache.Get(key)
				if !ok {
					t.Errorf("Key not found: %v", key)
				}

				if val != value {
					t.Errorf("Expected value: %v, got: %v", value, val)
				}
			}(i)
		}

		wg.Wait()
	})

	t.Run("Remove parallel", func(t *testing.T) {
		t.Parallel()
		var wg sync.WaitGroup

		for i := 0; i < 50; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				key := i
				cache.mu.Lock()
				defer cache.mu.Unlock()
				if node, ok := cache.items[key]; ok {
					cache.remove(node)
				}
			}()
		}

		wg.Wait()

		if cache.Len() != 0 {
			t.Errorf("Len failed")
		}
	})
}
