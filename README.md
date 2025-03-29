# ConcurrentQueue in Golang

## Overview

This project implements a thread-safe queue in Golang using a mutex for synchronization. The `ConcurrentQueue` struct allows multiple goroutines to enqueue and dequeue elements concurrently without data corruption.

## Features

- Thread-safe enqueue and dequeue operations
- Supports concurrent execution using `sync.Mutex`
- Provides a method to get the current size of the queue

## Installation

Clone the repository and navigate to the project directory:

```sh
git clone https://github.com/vaidik-bajpai/concurrent-queue.git
cd concurrent-queue
```

## Usage

Run the main program:

```sh
go run main.go
```

This will enqueue some elements, dequeue them, and print the results.

## API

### `Enqueue(item int32)`

Adds an item to the queue.

### `Dequeue() int32`

Removes and returns the front item from the queue. Panics if the queue is empty.

### `Size() int`

Returns the number of elements currently in the queue.

## Running Tests

The project includes a test to verify the concurrency safety of the queue. Run the tests using:

```sh
go test -v
```

This test enqueues 1,000,000 elements concurrently and verifies that the queue size matches the expected count.

## Example

```go
q := ConcurrentQueue{
	queue: make([]int32, 0),
}
q.Enqueue(10)
q.Enqueue(20)
fmt.Println(q.Dequeue()) // Output: 10
fmt.Println(q.Size())    // Output: 1
```
