package main

import (
	"fmt"
	"log"
	"lru/lru_pool"
	"time"
)

func main() {
	capacity := 8192
	cache := lru_pool.NewCache(capacity)

	cache.Add("key1", "value1")

	cache.AddWithTTL("key2", "value2", time.Minute)

	value, ok := cache.Get("key1")
	if !ok {
		log.Fatal("unexcisting key")
	}
	fmt.Println(value)

	value, ok = cache.Get("key2")
	if !ok {
		log.Fatal("unexcisting key")
	}
	fmt.Println(value)
}
