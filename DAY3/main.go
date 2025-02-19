package main

import (
    "errors"
    "fmt"
)

// Check if a number is prime
func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

// Divide two numbers with error handling
func safeDivide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    // Loop to print numbers from 1 to 20
    for i := 1; i <= 20; i++ {
        fmt.Print(i, " ")
    }
    fmt.Println()

    // Check if a number is prime
    num := 17
    if isPrime(num) {
        fmt.Printf("%d is a prime number\n", num)
    } else {
        fmt.Printf("%d is not a prime number\n", num)
    }

    // Test division with error handling
    result, err := safeDivide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Division Result:", result)
    }
}
