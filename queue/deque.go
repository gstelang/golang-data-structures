package main

type DoubleEndedQueue[T any] struct {
	head  int
	tail  int
	items []T
}

// A  -> B -> C -> D
// ^head
// A1 -> A -> B -> C -> D
// ^head
// Impacts the queue operation
// it's like a skipping a line if you're a queue
func (h *DoubleEndedQueue[T]) pushFront(item T) {
	h.items = append(h.items)
}

// A    -> B -> C -> D
// Head^
// B    -> C -> D
// Head^
// queue operation
func (h *DoubleEndedQueue[T]) popFront(item T) {

}

// A -> B -> C -> D
//
//	^ tail
//
// A -> B -> C -> D -> E
//
//	^ tail
//
// Stack & Queue operation
func (h *DoubleEndedQueue[T]) pushBack(item T) {

}

// A -> B -> C -> D -> E
//
// A -> B -> C -> D
// stack operation
func (h *DoubleEndedQueue[T]) popBack(item T) {

}

func growIfFull() {

}

func resize() {

}

func shrinkIfExcess() {

}

func NewDoubleEndedQueue[T any]() *DoubleEndedQueue[T] {
	return &DoubleEndedQueue[T]{
		head:  0,
		tail:  0,
		items: make([]T, 0),
	}
}

func main() {

}
