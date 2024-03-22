package main

import "fmt"

type Queue[T any] struct {
	items []T
}

func (s *Queue[T]) enqueue(item T) {
	s.items = append(s.items, item)
}

func (s *Queue[T]) dequeue() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	val := s.items[0]
	s.items = s.items[1:len(s.items)]
	return val, true
}

func (s *Queue[T]) size() int {
	return len(s.items)
}

func (s *Queue[T]) peek() T {
	return s.items[0]
}

func main() {
	seq := Queue[string]{}
	val0, check0 := seq.dequeue()
	fmt.Println(val0)
	fmt.Println(check0)
	seq.enqueue("abc")
	seq.enqueue("xyz")
	fmt.Println("Queue size:", seq.size())
	fmt.Println("Queue peek:", seq.peek())
	val, check := seq.dequeue()
	fmt.Println("Queue peek:", seq.peek())
	fmt.Println(val)
	fmt.Println(check)
	fmt.Println("Queue size:", seq.size())

	seqInts := Queue[int]{}
	valInt, checkInt := seqInts.dequeue()
	fmt.Println(valInt)
	fmt.Println(checkInt)
	seqInts.enqueue(1)
	seqInts.enqueue(2)
	fmt.Println("Queue size:", seqInts.size())
	fmt.Println("Queue peek:", seqInts.peek())
	valIntPopped, checkPopped := seqInts.dequeue()
	fmt.Println(valIntPopped)
	fmt.Println(checkPopped)
	fmt.Println("Queue size:", seq.size())

}
