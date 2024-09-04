package lru

// func TestPriorityQueueWithTTL1(t *testing.T) {
// 	capacity := 3
// 	c := NewCache(capacity)

// 	c.AddWithTTL(1, "value1", 4*time.Hour)
// 	c.AddWithTTL(2, "value2", 3*time.Hour)
// 	c.AddWithTTL(3, "value3", 0)

// 	assert.Equal(t, 3, c.pq.Len(), "Queue length should be 3")

// 	_, _ = c.Get(1)
// 	assert.Equal(t, 3, c.pq.Len(), "Queue length should be 3")

// 	c.AddWithTTL(5, "value1", 2*time.Hour)
// 	assert.Equal(t, 3, c.pq.Len(), "Queue length should be 3")
// 	c.Add(5, "value3")

// 	assert.Equal(t, 3, c.pq.Len(), "Queue length should be 3")
// }

// func TestPriorityQueueWithTTL12(t *testing.T) {
// 	capacity := 3
// 	c := NewCache(capacity)

// 	c.AddWithTTL(1, "value1", 2*time.Second) // TTL = 2s
// 	c.AddWithTTL(2, "value2", 1*time.Second) // TTL = 1s
// 	c.AddWithTTL(3, "value3", 0)             // TTL = 0 (без истечения)

// 	c.AddWithTTL(3, "value3", 0)
// 	c.AddWithTTL(2, "value3", 0)             // Должен переместиться на конец
// 	c.AddWithTTL(5, "value1", 2*time.Second) // Добавление нового элемента

// 	assert.Equal(t, 3, c.pq.Len(), "Queue length should be 3")

// }

// func TestOrderAddPop(t *testing.T) {
// 	// смотрим порядок при Pop
// 	capacity := 3
// 	c := NewCache(capacity)

// 	c.Add("key1", "")
// 	c.Add("key2", "")
// 	c.Add("key3", "")

// 	assert.Equal(t, 3, c.pq.Len(), "")

// 	node1 := c.pq.Pop()
// 	assert.Equal(t, "key1", node1.key, "")

// 	node2 := c.pq.Pop()
// 	assert.Equal(t, "key2", node2.key, "")

// 	node3 := c.pq.Pop()
// 	assert.Equal(t, "key3", node3.key, "")

// 	assert.Equal(t, 0, c.pq.Len(), "")
// }

// func TestOrderAddWithTTLPop(t *testing.T) {
// 	// смотрим порядок при Pop
// 	capacity := 3
// 	c := NewCache(capacity)

// 	c.AddWithTTL("key1", "value1", 2*time.Second) // TTL = 2s
// 	c.AddWithTTL("key2", "value2", 1*time.Second) // TTL = 1s
// 	c.AddWithTTL("key3", "value3", 0)             // TTL = 0 (без истечения)

// 	assert.Equal(t, 3, c.pq.Len(), "")

// 	node1 := c.pq.Pop()
// 	assert.Equal(t, "key3", node1.key, "")

// 	node2 := c.pq.Pop()
// 	assert.Equal(t, "key2", node2.key, "")

// 	node3 := c.pq.Pop()
// 	assert.Equal(t, "key1", node3.key, "")

// 	assert.Equal(t, 0, c.pq.Len(), "")
// }
