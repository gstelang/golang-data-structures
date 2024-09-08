package main

import "fmt"

type RingBuffer struct {
	logs         []int
	totalEntries int
	capacity     int
}

type RingOperations interface {
	Add(int)
	GetRecentEntries() []int
}

func (rb *RingBuffer) Add(val int) {
	nextIndex := rb.totalEntries % rb.capacity
	rb.logs[nextIndex] = val
	rb.totalEntries++
}

func NewRingBuffer(capacity int) *RingBuffer {
	return &RingBuffer{
		logs:         make([]int, capacity),
		capacity:     capacity,
		totalEntries: 0,
	}
}

func (rb *RingBuffer) GetRecentEntries() []int {

	index := rb.totalEntries % rb.capacity

	result := make([]int, rb.capacity)
	resultIndex := 0
	for i := index; i < rb.capacity; i++ {
		result[resultIndex] = rb.logs[i]
		resultIndex++
	}

	for j := 0; j < index; j++ {
		result[resultIndex] = rb.logs[j]
		resultIndex++
	}

	return result
}

func main() {

	ringBuffer := NewRingBuffer(5)

	ringBuffer.Add(1)
	ringBuffer.Add(2)
	ringBuffer.Add(3)
	ringBuffer.Add(4)
	ringBuffer.Add(5)
	ringBuffer.Add(6)
	ringBuffer.Add(7)

	recent := ringBuffer.GetRecentEntries()
	fmt.Println(recent)
}
