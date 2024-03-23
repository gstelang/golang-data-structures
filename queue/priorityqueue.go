package main

import (
	"container/heap"
	"fmt"

	"golang.org/x/exp/constraints"
)

// sort.Interface
// type Interface interface {
// 	Len() int
// 	Less(i, j int) bool
// 	Swap(i, j int)
// }

// heap.Interface
// type Interface interface {
// 	sort.Interface
// 	Push(x interface{}) // add x as element Len()
// 	Pop() interface{}   // remove and return element Len() - 1.
// }

// Heap interface is composed of sort interface.

type PriorityQueue[T constraints.Ordered] struct {
	items    []T
	heapType string
}

func (h *PriorityQueue[T]) Len() int {
	return len(h.items)
}

func (h *PriorityQueue[T]) Less(i, j int) bool {
	defaultMin := h.items[i] < h.items[j]
	if h.heapType == "MAX" {
		return h.items[i] > h.items[j]
	}

	return defaultMin
}

func (h *PriorityQueue[T]) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *PriorityQueue[T]) Push(item interface{}) {
	h.items = append(h.items, item.(T))
}

// This is the default implementation

// Swaps the first element to the end
// Fix the heap
// Call your Pop() implementation

// func Pop(h Interface) any {
// 	n := h.Len() - 1
// 	h.Swap(0, n)
// 	down(h, 0, n)
// 	return h.Pop()
// }

func (h *PriorityQueue[T]) Pop() interface{} {
	if len(h.items) == 0 {
		fmt.Println("Heap is empty. Cannot pop.")
		return nil
	}
	old := h.items
	n := len(old)

	lastElem := old[n-1]
	// slice exlcuding the last elem.
	h.items = old[0 : n-1]
	return lastElem
}

func main() {
	pq := &PriorityQueue[int]{
		heapType: "MIN",
	}
	heap.Init(pq)
	heap.Push(pq, 9)
	heap.Push(pq, 20)
	heap.Push(pq, 8)
	heap.Push(pq, 3)
	heap.Push(pq, 100)
	heap.Push(pq, 99)
	heap.Push(pq, 40)

	fmt.Println(pq)
	for pq.Len() > 0 {
		x := heap.Pop(pq)
		fmt.Println(x)
	}

	pqMax := &PriorityQueue[int]{
		heapType: "MAX",
	}
	heap.Init(pqMax)
	heap.Push(pqMax, 9)
	heap.Push(pqMax, 20)
	heap.Push(pqMax, 8)
	heap.Push(pqMax, 3)
	heap.Push(pqMax, 100)
	heap.Push(pqMax, 99)
	heap.Push(pqMax, 40)

	fmt.Println(pqMax)
	for pqMax.Len() > 0 {
		x := heap.Pop(pqMax)
		fmt.Println(x)
	}
}
