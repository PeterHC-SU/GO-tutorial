package main

import "fmt"

func main() {
    // Array of 5 numbers
    numbers := [5]int{1, 2, 3, 4, 5}
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    fmt.Println("Sum of array elements:", sum)

    // Slice of names
    names := []string{"Alice", "Bob", "Charlie"}
    names = append(names, "David") // Adding a new name
    fmt.Println("Names:", names)

    // Map of countries and their capitals
    capitals := map[string]string{
        "USA":    "Washington D.C.",
        "France": "Paris",
        "Japan":  "Tokyo",
    }
    fmt.Println("Country Capitals:", capitals)
}
