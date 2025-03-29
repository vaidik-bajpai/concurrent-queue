package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcurrenySafety(t *testing.T) {
	q := ConcurrentQueue{
		queue: make([]int32, 0),
	}

	var wg sync.WaitGroup
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			q.Enqueue(rand.Int31())
			wg.Done()
		}()
	}
	wg.Wait()

	assert.Equal(t, 1000000, q.Size(), fmt.Sprintf("The expected size [%d] and actual size [%d] of the queue don't match", 1000000, q.Size()))
}
