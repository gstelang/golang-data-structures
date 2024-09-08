package main

import "fmt"

type RingBuffer[T any] struct {
	logs         []T
	totalEntries int
	capacity     int
}

type RingOperations[T any] interface {
	Add(T)
	GetRecentEntries() []T
}

func (rb *RingBuffer[T]) Add(val T) {
	nextIndex := rb.totalEntries % rb.capacity
	rb.logs[nextIndex] = val
	rb.totalEntries++
}

func NewRingBuffer[T any](capacity int) *RingBuffer[T] {
	return &RingBuffer[T]{
		logs:         make([]T, capacity),
		capacity:     capacity,
		totalEntries: 0,
	}
}

func (rb *RingBuffer[T]) GetRecentEntries() []T {

	size := min(rb.totalEntries, rb.capacity)
	result := make([]T, size)

	index := rb.totalEntries % rb.capacity
	resultIndex := 0

	for j := 0; j < index; j++ {
		result[resultIndex] = rb.logs[j]
		resultIndex++
	}

	if rb.totalEntries >= rb.capacity {
		for i := index; i < rb.capacity; i++ {
			result[resultIndex] = rb.logs[i]
			resultIndex++
		}
	}

	return result
}

func main() {

	// int
	ringBufferInt := NewRingBuffer[int](5)
	ringBufferInt.Add(1)
	ringBufferInt.Add(2)
	ringBufferInt.Add(3)
	ringBufferInt.Add(4)
	ringBufferInt.Add(5)
	ringBufferInt.Add(6)
	ringBufferInt.Add(7)
	recent := ringBufferInt.GetRecentEntries()
	fmt.Println(recent)

	// str
	ringBufferStr := NewRingBuffer[string](5)
	ringBufferStr.Add("foo")
	ringBufferStr.Add("bar")
	ringBufferStr.Add("baz")
	recentStr := ringBufferStr.GetRecentEntries()
	fmt.Println(recentStr)

	// bytes
	byteBuffer := NewRingBuffer[[]byte](5)
	byteBuffer.Add([]byte("ABC"))
	byteBuffer.Add([]byte("DEF"))
	byteBuffer.Add([]byte("GHI"))
	byteBuffer.Add([]byte("JKL"))
	byteBuffer.Add([]byte("MNO"))
	byteBuffer.Add([]byte("PQR"))
	byteBuffer.Add([]byte("STU"))

	for _, bytes := range byteBuffer.GetRecentEntries() {
		fmt.Println(string(bytes))
	}
}
