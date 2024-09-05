package lru_pool

import (
	"sync"
	"testing"
	"time"
)

func TestCacheRaceConditionParallelOnlyOneGet(t *testing.T) {
	cache := NewCache(1)

	t.Run("Add and Get parallel", func(t *testing.T) {
		t.Parallel()
		var wg sync.WaitGroup

		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				_, ok := cache.Get(1)
				if ok {
					t.Errorf("Key not found: %v", i)
				}
			}()
		}
		wg.Wait()
	})
}

func TestCacheRaceConditionParallelOnlyOneGetAdd(t *testing.T) {
	cache := NewCache(1)

	t.Run("Add and Get parallel", func(t *testing.T) {
		t.Parallel()
		var wg sync.WaitGroup

		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				value := 1
				cache.AddWithTTL(1, value, time.Minute)

				val, ok := cache.Get(1)
				if !ok {
					t.Errorf("Key not found: %v", i)
				}
				if val != value {
					t.Errorf("Expected value: %v, got: %v", value, val)
				}
			}()
		}
		wg.Wait()
	})
}

func TestCacheRaceConditionParallelAddingAndParrallelRemove(t *testing.T) {
	cache := NewCache(100)
	t.Run("Add and Get parallel", func(t *testing.T) {
		t.Parallel()

		var wg sync.WaitGroup

		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				value := i * 10
				cache.AddWithTTL(i, value, time.Minute)
			}()
		}
		wg.Wait()

	})

	t.Run("Remove parallel", func(t *testing.T) {
		t.Parallel()
		// не совсем уверен в корректности удаления в параллельном запуске
		var wg sync.WaitGroup

		for i := 0; i < 100; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()
				cache.Remove(i)
			}()
		}
		wg.Wait()
	})

	if cache.Len() != 0 {
		t.Errorf("Cache size should be 0, but got %d", cache.Len())
	}
}
