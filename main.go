package main

import (
	"container/heap"
)

func main() {
	// deq := newDeque{}
	pq := &intHeap{}
	sc := make([]int, 4)
	uf := UnionFind{sc}
	heap.Init(pq)
}
