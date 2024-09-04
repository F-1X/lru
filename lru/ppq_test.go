package lru

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRemove3(t *testing.T) {
	pqq := NewPQQ()

	node1 := &Node{key: 1, value: "", expiration: time.Now().Add(2 * time.Hour).UnixNano()}
	node2 := &Node{key: 2, value: ""}
	node3 := &Node{key: 3, value: "", expiration: time.Now().Add(3 * time.Hour).UnixNano()}
	node4 := &Node{key: 4, value: "", expiration: time.Now().Add(3 * time.Hour).UnixNano()}

	pqq.Push(node1)
	pqq.Push(node2)
	pqq.Push(node3)
	pqq.Push(node4)

	assert.Equal(t, 4, pqq.Len(), "Len != 3")

	pqq.Remove(node1)

	first := pqq.Pop()
	assert.Equal(t, 2, first.key, "order 1")

	// 3 4 5
	node5 := &Node{key: 5, value: "", expiration: time.Now().Add(3 * time.Hour).UnixNano()}
	pqq.Push(node5)

	second := pqq.Pop()
	third := pqq.Pop()
	four := pqq.Pop()

	assert.Equal(t, 3, second.key, "order 2")
	assert.Equal(t, 4, third.key, "order 3")
	assert.Equal(t, 5, four.key, "order 4")

	assert.Equal(t, 0, pqq.Len(), "Len != 0")
}

func TestRemove2(t *testing.T) {
	pqq := NewPQQ()

	node1 := &Node{key: 1, value: "", expiration: time.Now().Add(2 * time.Hour).UnixNano()}
	node2 := &Node{key: 2, value: ""}
	node3 := &Node{key: 3, value: "", expiration: time.Now().Add(3 * time.Hour).UnixNano()}
	node4 := &Node{key: 4, value: "", expiration: time.Now().Add(3 * time.Hour).UnixNano()}

	pqq.Push(node1)
	pqq.Push(node2)
	pqq.Push(node3)
	pqq.Push(node4)

	assert.Equal(t, 4, pqq.Len(), "Len != 3")

	pqq.Remove(node1)

	first := pqq.Pop()
	assert.Equal(t, 2, first.key, "order 1")

	// 3 4 5
	node5 := &Node{key: 5, value: "", expiration: time.Now().Add(3 * time.Hour).UnixNano()}
	pqq.Push(node5)

	second := pqq.Pop()
	third := pqq.Pop()
	four := pqq.Pop()

	assert.Equal(t, 3, second.key, "order 2")
	assert.Equal(t, 4, third.key, "order 3")
	assert.Equal(t, 5, four.key, "order 4")

	assert.Equal(t, 0, pqq.Len(), "Len != 0")
}

func TestRemove1(t *testing.T) {
	pqq := NewPQQ()

	node1 := &Node{key: 1, value: "", expiration: time.Now().Add(2 * time.Hour).UnixNano()}
	node2 := &Node{key: 2, value: ""}
	node3 := &Node{key: 3, value: "", expiration: time.Now().Add(3 * time.Hour).UnixNano()}

	pqq.Push(node1)
	pqq.Push(node2)
	pqq.Push(node3)

	assert.Equal(t, 3, pqq.Len(), "Len != 3")

	pqq.Remove(node2)

	first := pqq.Pop()
	second := pqq.Pop()

	assert.Equal(t, 1, first.key, "order 1")
	assert.Equal(t, 3, second.key, "order 2")

	assert.Equal(t, 0, pqq.Len(), "Len != 0")
}

func TestPriorityQueue(t *testing.T) {
	pqq := NewPQQ()

	pqq.Push(&Node{key: 1, value: "", expiration: time.Now().Add(1 * time.Hour).UnixNano()})
	pqq.Push(&Node{key: 2, value: "", expiration: time.Now().Add(1 * time.Hour).UnixNano()})
	pqq.Push(&Node{key: 3, value: "", expiration: time.Now().Add(1 * time.Hour).UnixNano()})

	assert.Equal(t, 3, pqq.Len(), "Len != 3")

	first := pqq.Pop()
	second := pqq.Pop()
	third := pqq.Pop()

	assert.Equal(t, 1, first.key, "order 1")
	assert.Equal(t, 2, second.key, "order 2")
	assert.Equal(t, 3, third.key, "order 3")

	assert.Equal(t, 0, pqq.Len(), "Len != 0")
}

func TestPriorityQueueDifferenctTime(t *testing.T) {
	pqq := NewPQQ()

	pqq.Push(&Node{key: 1, value: "", expiration: time.Now().Add(2 * time.Hour).UnixNano()})
	pqq.Push(&Node{key: 2, value: "", expiration: time.Now().Add(1 * time.Hour).UnixNano()})
	pqq.Push(&Node{key: 3, value: "", expiration: time.Now().Add(3 * time.Hour).UnixNano()})

	assert.Equal(t, 3, pqq.Len(), "Len != 3")

	first := pqq.Pop()
	second := pqq.Pop()
	third := pqq.Pop()

	assert.Equal(t, 2, first.key, "order 1")
	assert.Equal(t, 1, second.key, "order 2")
	assert.Equal(t, 3, third.key, "order 3")

	assert.Equal(t, 0, pqq.Len(), "Len != 0")
}

func TestPriorityQueueMixed(t *testing.T) {
	pqq := NewPQQ()

	pqq.Push(&Node{key: 1, value: "", expiration: time.Now().Add(2 * time.Hour).UnixNano()})
	pqq.Push(&Node{key: 2, value: ""})
	pqq.Push(&Node{key: 3, value: "", expiration: time.Now().Add(3 * time.Hour).UnixNano()})

	assert.Equal(t, 3, pqq.Len(), "Len != 3")

	first := pqq.Pop()
	second := pqq.Pop()
	third := pqq.Pop()

	assert.Equal(t, 2, first.key, "order 1")
	assert.Equal(t, 1, second.key, "order 2")
	assert.Equal(t, 3, third.key, "order 3")

	assert.Equal(t, 0, pqq.Len(), "Len != 0")
}

func TestPriorityAddMixedTimes(t *testing.T) {
	pqq := NewPQQ()

	pqq.Push(&Node{key: 1, value: "", expiration: time.Now().Add(2 * time.Hour).UnixNano()})
	pqq.Push(&Node{key: 2, value: ""})
	pqq.Push(&Node{key: 3, value: "", expiration: time.Now().Add(3 * time.Hour).UnixNano()})

	assert.Equal(t, 3, pqq.Len(), "Len != 3")

	first := pqq.Pop()
	second := pqq.Pop()
	third := pqq.Pop()
	assert.Equal(t, 2, first.key, "order 1")
	assert.Equal(t, 1, second.key, "order 2")
	assert.Equal(t, 3, third.key, "order 3")

	assert.Equal(t, 0, pqq.Len(), "Len != 0")
}
