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

	size := min(rb.totalEntries, rb.capacity)
	result := make([]int, size)

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
