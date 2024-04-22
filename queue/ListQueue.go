package queue

import (
	"container/list"
)

type ListQueue[T any] struct {
	items list.List
}

func (s *ListQueue[T]) enqueue(item T) {
	s.items.PushBack(item)
}

func (s *ListQueue[T]) dequeue() (any, bool) {
	if s.items.Len() == 0 {
		var zero T
		return zero, false
	}
	elem := s.items.Front()
	s.items.Remove(elem)
	return elem.Value, true
}

func (s *ListQueue[T]) size() int {
	return s.items.Len()
}

// any is an alias for interface{} and is equivalent to interface{} in all ways.
// type any = interface{}
func (s *ListQueue[T]) peek() any {
	return s.items.Front().Value
}

// func main() {
// 	seq := ListQueue[string]{}
// 	val0, check0 := seq.dequeue()
// 	fmt.Println(val0)
// 	fmt.Println(check0)
// 	seq.enqueue("abc")
// 	seq.enqueue("xyz")
// 	fmt.Println("ListQueue size:", seq.size())
// 	fmt.Println("ListQueue peek:", seq.peek())
// 	val, check := seq.dequeue()
// 	fmt.Println("ListQueue peek:", seq.peek())
// 	fmt.Println(val)
// 	fmt.Println(check)
// 	fmt.Println("ListQueue size:", seq.size())

// 	seqInts := ListQueue[int]{}
// 	valInt, checkInt := seqInts.dequeue()
// 	fmt.Println(valInt)
// 	fmt.Println(checkInt)
// 	seqInts.enqueue(1)
// 	seqInts.enqueue(2)
// 	fmt.Println("Queue size:", seqInts.size())
// 	fmt.Println("Queue peek:", seqInts.peek())
// 	valIntPopped, checkPopped := seqInts.dequeue()
// 	fmt.Println(valIntPopped)
// 	fmt.Println(checkPopped)
// 	fmt.Println("Queue size:", seq.size())
// }
