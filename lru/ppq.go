package lru

import (
	"container/heap"
	"log"
)

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Сначала сортируем по expiration (по возрастанию)
	// Если expiration равен, то сортируем по index (по возрастанию)
	log.Println("pq[i].expiration", pq[i].expiration, "pq[j].expiration", pq[j].expiration)
	if pq[i].expiration == 0 && pq[j].expiration != 0 {
		log.Println("hit 1")
		return true
	}
	if pq[i].expiration != 0 && pq[j].expiration == 0 {
		log.Println("hit 2")
		return false
	}
	if pq[i].expiration == pq[j].expiration {
		log.Println("hit 3", "key:", pq[i].key, pq[i].index, "key:", pq[j].key, pq[j].index)
		// return pq[i].index < pq[j].index
	}
	log.Println("hit 4")
	return pq[i].expiration < pq[j].expiration
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.index = -1
	*pq = old[0 : n-1]

	// old := *pq
	// n := len(old)
	// node := old[0]
	// old[0] = nil
	// node.index = -1
	// *pq = old[1:n]
	return node
}

func (pq *PriorityQueue) Remove(node *Node) {
	if node.index >= 0 && node.index < len(*pq) {
		heap.Remove(pq, node.index)
		node.index = -1
	}
}

type PQQ struct {
	pq *PriorityQueue
}

func NewPQQ() *PQQ {
	pq := &PriorityQueue{}
	heap.Init(pq)
	return &PQQ{pq: pq}
}

func (q *PQQ) Len() int {
	return len(*q.pq)
}

func (q *PQQ) Push(node *Node) {
	heap.Push(q.pq, node)
}

func (q *PQQ) Pop() *Node {
	if q.Len() == 0 {
		return nil
	}
	return heap.Pop(q.pq).(*Node)
}

func (q *PQQ) Remove(node *Node) {
	q.pq.Remove(node)
}
