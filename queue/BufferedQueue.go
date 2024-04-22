// just a fun thought exercise with no practical application.
// Better to use slice for small use cases or a list
package queue

import (
	"fmt"
)

type BufferedQueue[T any] struct {
	bufferedItems chan T
	bufferSize    int
	count         int
}

func (s *BufferedQueue[T]) enqueue(item T) {
	if s.count == s.bufferSize {
		fmt.Println("Buffer is full. Cannot add more items")
		return
	}
	if s.count == 0 {
		s.bufferedItems = make(chan T, s.bufferSize)
	}
	s.bufferedItems <- item
	s.count++
}

func (s *BufferedQueue[T]) dequeue() (any, bool) {
	if s.count == 0 {
		var zero T
		return zero, false
	}
	val := <-s.bufferedItems
	s.count--
	return val, true
}

func (s *BufferedQueue[T]) size() int {
	return s.count
}

// peek is not possible
// func (s *BufferedQueue[T]) peek() any {}

// func main() {
// 	seq := BufferedQueue[string]{bufferSize: 2}
// 	val0, check0 := seq.dequeue()
// 	fmt.Println(val0)
// 	fmt.Println(check0)
// 	seq.enqueue("abc")
// 	seq.enqueue("xyz")
// 	seq.enqueue("efg")
// 	fmt.Println("BufferedQueue size:", seq.size())
// 	val, check := seq.dequeue()
// 	fmt.Println(val)
// 	fmt.Println(check)
// 	fmt.Println("BufferedQueue size:", seq.size())

// 	seqInts := BufferedQueue[int]{bufferSize: 20}
// 	valInt, checkInt := seqInts.dequeue()
// 	fmt.Println(valInt)
// 	fmt.Println(checkInt)
// 	seqInts.enqueue(1)
// 	seqInts.enqueue(2)
// 	fmt.Println("BufferedQueue size:", seqInts.size())
// 	// fmt.Println("Queue peek:", seqInts.peek())
// 	valIntPopped, checkPopped := seqInts.dequeue()
// 	fmt.Println(valIntPopped)
// 	fmt.Println(checkPopped)
// 	fmt.Println("BufferedQueue size:", seq.size())

// }
