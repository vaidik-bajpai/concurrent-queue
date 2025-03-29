package main

import (
	"fmt"
	"sync"
)

type ConcurrentQueue struct {
	queue []int32
	mu    sync.Mutex
}

func (q *ConcurrentQueue) Enqueue(item int32) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = append(q.queue, item)
}

func (q *ConcurrentQueue) Dequeue() int32 {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.queue) == 0 {
		panic("[underflow condition] cannot dequeue from an empty queue")
	}

	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func (q *ConcurrentQueue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.queue)
}

func main() {
	q := ConcurrentQueue{
		queue: make([]int32, 0),
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
