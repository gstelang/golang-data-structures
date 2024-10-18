package main

import (
	"fmt"
	"log"
)

type CircularQueue struct {
	currentSize int
	capacity    int
	front       int
	rear        int
	arr         []int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		currentSize: 0,
		capacity:    capacity,
		front:       -1,
		rear:        -1,
		arr:         make([]int, capacity),
	}
}

// Test cases
// 1. enqueue and then dequeue
// 2. enqueue, enqueue, dequeue
// 3. enqueue, enqueue, enqueque, enqueue, enqueue (error), dequeue
// 4. dequeue (error)
// 5. enqueue

func (cq *CircularQueue) enqueue(item int) {
	// Handle when queue is full
	if cq.currentSize == cq.capacity { //
		fmt.Println("Circular queue is full")
		return
	}

	// enqueue just changes the rear.
	cq.arr = append(cq.arr, item)

	if cq.front == -1 {
		cq.front = cq.front + 1 // 0
	}
	cq.rear = cq.rear + 1               // 0
	cq.currentSize = cq.currentSize + 1 // 1
}

func (cq *CircularQueue) dequeue(item int) (int, error) {
	// Handle when queue is empty
	if cq.currentSize == 0 {
		log.Fatalf("Cannot remove from a empty queue")
	}

	val := cq.arr[cq.front] // 0
	if cq.rear > cq.front {
		cq.front = cq.front + 1
	}
	cq.currentSize = cq.currentSize - 1
	return val, nil // 0
}

// func (cq *CircularQueue) isFull(item int) bool {

// }

// func (cq *CircularQueue) isEmpty(item int) bool {

// }

func main() {
	cq := NewCircularQueue(4)
	fmt.Println(cq)
}
