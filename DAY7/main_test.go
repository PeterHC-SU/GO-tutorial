package main

import (
	"sync"
	"testing"
	"time"
)

// TEST printNumbers
func TestPrintNumbers(t *testing.T) {
	go printNumbers(1)
	go printNumbers(2)

	time.Sleep(time.Second) // 等待 goroutine 執行
}

// TEST Worker Pool
func TestWorkerPool(t *testing.T) {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup

	// 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Sent 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait() // make sure worker done
	close(results)

	expectedResults := []int{2, 4, 6, 8, 10}
	var actualResults []int

	for result := range results {
		actualResults = append(actualResults, result)
	}

	if len(actualResults) != len(expectedResults) {
		t.Errorf("Worker pool failed: expected %d results, got %d", len(expectedResults), len(actualResults))
	}

	// for i, v := range actualResults {
	// 	if v != expectedResults[i] {
	// 		t.Errorf("Worker pool output mismatch at index %d: expected %d, got %d", i, expectedResults[i], v)
	// 	}
	// }
}

// TEST Goroutine message
func TestMessageChannel(t *testing.T) {
	messages := make(chan string)

	go func() {
		messages <- "Hello from goroutine!"
	}()

	received := <-messages
	expected := "Hello from goroutine!"

	if received != expected {
		t.Errorf("Channel message failed: expected %s, got %s", expected, received)
	}
}
