package queue

type Queue[T any] struct {
	items []T
}

func (s *Queue[T]) Enqueue(item T) {
	s.items = append(s.items, item)
}

func (s *Queue[T]) Dequeue() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	val := s.items[0]
	s.items = s.items[1:len(s.items)]
	return val, true
}

func (s *Queue[T]) Size() int {
	return len(s.items)
}

func (s *Queue[T]) Peek() T {
	return s.items[0]
}

// func main() {
// 	seq := Queue[string]{}
// 	val0, check0 := seq.Dequeue()
// 	fmt.Println(val0)
// 	fmt.Println(check0)
// 	seq.Enqueue("abc")
// 	seq.Enqueue("xyz")
// 	fmt.Println("Queue Size:", seq.Size())
// 	fmt.Println("Queue Peek:", seq.Peek())
// 	val, check := seq.Dequeue()
// 	fmt.Println("Queue Peek:", seq.Peek())
// 	fmt.Println(val)
// 	fmt.Println(check)
// 	fmt.Println("Queue Size:", seq.Size())

// 	seqInts := Queue[int]{}
// 	valInt, checkInt := seqInts.Dequeue()
// 	fmt.Println(valInt)
// 	fmt.Println(checkInt)
// 	seqInts.Enqueue(1)
// 	seqInts.Enqueue(2)
// 	fmt.Println("Queue Size:", seqInts.Size())
// 	fmt.Println("Queue Peek:", seqInts.Peek())
// 	valIntPopped, checkPopped := seqInts.Dequeue()
// 	fmt.Println(valIntPopped)
// 	fmt.Println(checkPopped)
// 	fmt.Println("Queue Size:", seq.Size())

// }
