package main

import "fmt"

// Function that returns the sum of two numbers
func add(a, b int) int {
    return a + b
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero is not allowed")
    }
    return a / b, nil
}

func main() {
    // Variable declaration
    var message string = "Hello, Go Variables!"
    number := 42 // Short declaration
    const pi float64 = 3.14159

    fmt.Println(message)
    fmt.Println("Number:", number)
    fmt.Println("Pi:", pi)

    // Function usage
    sum := add(10, 20)
    fmt.Println("Sum:", sum)

    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Division Result:", result)
    }
}
