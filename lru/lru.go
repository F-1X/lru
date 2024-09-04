package lru

import (
	"container/heap"
	"log"
	"sync"
	"time"
)

type Node struct {
	key        any
	value      any
	expiration int64 // 0 означает без истечения
	index      int   // индекс - позиция в приоритетном списке
}

type Cache struct {
	cap   int
	cache map[any]*Node
	len   int
	mu    sync.Mutex
	pq    *PQQ // Используем приоритетную очередь для управления элементами
}

func NewCache(cap int) *Cache {
	return &Cache{
		cap:   cap,
		cache: make(map[any]*Node, cap),
		pq:    NewPQQ(),
		mu:    sync.Mutex{},
	}
}

// func (c *Cache) Get(key any) (any, bool) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	node, found := c.cache[key]
// 	if !found {
// 		return nil, false
// 	}

// 	// Проверка на истечение TTL
// 	if node.expiration != 0 && time.Now().UnixNano() > node.expiration {
// 		c.remove(node)
// 		return nil, false
// 	}

// 	// Обновляем приоритет в очереди
// 	c.pq.Remove(node)
// 	// c.pq.Push(node)
// 	// heap.Remove(c.pq.pq, node.index)
// 	heap.Push(c.pq.pq, node)
// 	return node.value, true
// }

func (c *Cache) Add(key, value any) {
	c.AddWithTTL(key, value, 0)
}

// func (c *Cache) AddWithTTL(key, value any, ttl time.Duration) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	expiration := time.Now().Add(ttl).UnixNano()
// 	if ttl == 0 {
// 		expiration = 0
// 	}

// 	if node, found := c.cache[key]; found {
// 		if node.index == -1 {
// 			log.Println("Updating already evicted key:", node)
// 		}
// 		node.value = value
// 		node.expiration = expiration

// 		heap.Remove(c.pq.pq, node.index)
// 		heap.Push(c.pq.pq, node)
// 	} else {
// 		if c.len >= c.cap {
// 			c.evictWithTTL()
// 		}

// 		node := &Node{key: key, value: value, expiration: expiration}
// 		c.cache[key] = node

//			heap.Push(c.pq.pq, node)
//			c.len++
//		}
//	}
func (c *Cache) Get(key any) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	node, found := c.cache[key]
	if !found {
		return nil, false
	}

	// Проверка на истечение TTL
	if node.expiration != 0 && time.Now().UnixNano() > node.expiration {
		log.Println("GET:remove by expire", node)
		c.remove(node)
		return nil, false
	}

	// Обновляем приоритет в очереди
	heap.Remove(c.pq.pq, node.index)
	heap.Push(c.pq.pq, node)

	log.Println("GET:remove and pushin", node)
	return node.value, true
}
func (c *Cache) AddWithTTL(key, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	expiration := time.Now().Add(ttl).UnixNano()
	if ttl == 0 {
		expiration = 0
	}

	// Если элемент уже существует, обновляем его
	if node, found := c.cache[key]; found {
		if node.index == -1 {
			log.Println("Updating already evicted key:", node)
		}
		node.value = value
		node.expiration = expiration

		c.pq.Remove(node)
		heap.Push(c.pq.pq, node)

		log.Println("ADD:remove and pushin", node)
	} else {
		// Проверяем, нужно ли удалить старые элементы
		if c.len >= c.cap {
			log.Println("EVICT...", key)
			c.evictWithTTL()
		}
		node := &Node{key: key, value: value, expiration: expiration}
		c.cache[key] = node

		heap.Push(c.pq.pq, node)
		// c.pq.Push(node)

		log.Println("ADD:pushin", node)
		c.len++
	}
}

func (c *Cache) evictWithTTL() {
	for c.pq.Len() > 0 {
		node := heap.Pop(c.pq.pq).(*Node)
		log.Println("evicting: ", node.key, node.index)
		c.remove(node)
	}
}

func (c *Cache) remove(node *Node) {
	if node.index >= 0 {
		log.Println("remove: ", node.key)
		c.pq.Remove(node)
	}
	log.Println("delete key: ", node.key)
	delete(c.cache, node.key)
	c.len--
}

func (c *Cache) Cap() int {
	return c.cap
}

func (c *Cache) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.len
}

func (c *Cache) Clear() {
	panic("unimpl")
}
