package main

import (
    "fmt"
    "sync"
    "time"
)

// Function to run as a goroutine
func printNumbers(id int) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Goroutine %d: %d\n", id, i)
        time.Sleep(time.Millisecond * 500)
    }
}

// Worker function for worker pool
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(time.Second) // Simulating work
        results <- job * 2
    }
}

func main() {
    // Running two goroutines concurrently
    go printNumbers(1)
    go printNumbers(2)

    // Using a channel for communication
    messages := make(chan string)
    go func() {
        messages <- "Hello from goroutine!"
    }()
    fmt.Println(<-messages)

    // Worker pool with goroutines
    jobs := make(chan int, 5)
    results := make(chan int, 5)
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, jobs, results, &wg)
    }

    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    wg.Wait()
    close(results)

    for result := range results {
        fmt.Println("Result:", result)
    }
}
