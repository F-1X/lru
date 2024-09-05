package lru_pool

import (
	"sync"
	"time"
)

type Node struct {
	Key        any
	Value      any
	Expiration time.Time
	Prev       *Node
	Next       *Node
}

type List struct {
	Head *Node
	Tail *Node
	Size int
	pool *sync.Pool
}

func NewList() *List {
	return &List{
		pool: &sync.Pool{
			New: func() any {
				return &Node{}
			},
		},
	}
}

func (l *List) GetNode() *Node {
	return l.pool.Get().(*Node)
}

func (l *List) PutNode(node *Node) {
	node.Key = nil
	node.Value = nil
	node.Expiration = time.Time{}
	node.Prev = nil
	node.Next = nil
	l.pool.Put(node)
}

func (l *List) PushFront(key any, value any, expiration time.Time) *Node {
	node := l.GetNode()
	node.Key = key
	node.Value = value
	node.Expiration = expiration

	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head.Prev = node
		l.Head = node
	}

	l.Size++
	return node
}

func (l *List) Remove(node *Node) {
	if node == nil {
		return
	}

	if node == l.Head {
		l.Head = node.Next
	}

	if node == l.Tail {
		l.Tail = node.Prev
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	node.Prev = nil
	node.Next = nil

	l.Size--
	l.PutNode(node)
}

type Cache struct {
	capacity int
	items    map[any]*Node
	list     *List
	mu       *sync.Mutex
}

func NewCache(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		items:    make(map[any]*Node),
		list:     NewList(),
		mu:       &sync.Mutex{},
	}
}

func (c *Cache) Get(key any) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	node, ok := c.items[key]
	if !ok {
		return nil, false
	}

	if time.Now().After(node.Expiration) {
		c.remove(node)
		return nil, false
	}

	value := node.Value
	expiration := node.Expiration

	c.list.Remove(node)
	node = c.list.PushFront(key, value, expiration)
	c.items[key] = node
	return node.Value, true
}

func (c *Cache) Add(key any, value any) {
	c.AddWithTTL(key, value, time.Hour*24)
}

func (c *Cache) AddWithTTL(key any, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.items[key]; ok {
		c.list.Remove(node)
		node.Value = value
		node.Expiration = time.Now().Add(ttl)
		node = c.list.PushFront(key, value, node.Expiration)
		c.items[key] = node
		return
	}

	if c.list.Size == c.capacity {
		c.remove(c.list.Tail)
	}

	node := c.list.PushFront(key, value, time.Now().Add(ttl))
	c.items[key] = node
}

func (c *Cache) remove(node *Node) {
	if node == nil {
		return
	}
	delete(c.items, node.Key)
	c.list.Remove(node)
}

func (c *Cache) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.list.Size
}

func (c *Cache) Cap() int {
	return c.capacity
}

func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[any]*Node)
	c.list = NewList()
}

func (c *Cache) Remove(key any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	node, ok := c.items[key]
	if !ok || node == nil {
		return
	}

	delete(c.items, key)
	c.list.Remove(node)
}
