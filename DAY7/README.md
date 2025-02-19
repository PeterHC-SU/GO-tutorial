# Day 7: Concurrency & Goroutines (Optional)

## Learn:
- Goroutines (Basic Concurrency).
- Channels (Communication between Goroutines).

## Resources:
- [Goroutines](https://gobyexample.com/goroutines)
- [Channels](https://gobyexample.com/channels)

## Exercises:
1. **Run two goroutines that print numbers concurrently.**
2. **Use a channel to send and receive messages between goroutines.**
3. **Create a worker pool using goroutines.**

## Notes:
- Goroutines allow parallel execution using `go func()`.
- Channels provide safe communication between goroutines.
- The `sync.WaitGroup` helps manage multiple goroutines.
