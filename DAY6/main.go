package main

import (
    "fmt"
    "math"
)

// Function using pointers to swap two numbers
func swap(a, b *int) {
    *a, *b = *b, *a
}

// Interface Shape with Area() method
type Shape interface {
    Area() float64
}

// Circle struct
type Circle struct {
    Radius float64
}

// Rectangle struct
type Rectangle struct {
    Width, Height float64
}

// Implement Area() for Circle
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// Implement Area() for Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Struct with pointer receiver method
type Counter struct {
    Value int
}

func (c *Counter) Increment() {
    c.Value++
}

func main() {
    // Swap function using pointers
    x, y := 10, 20
    fmt.Println("Before Swap:", x, y)
    swap(&x, &y)
    fmt.Println("After Swap:", x, y)

    // Interface usage
    c := Circle{Radius: 5}
    r := Rectangle{Width: 4, Height: 6}
    fmt.Println("Circle Area:", c.Area())
    fmt.Println("Rectangle Area:", r.Area())

    // Pointer receiver method
    counter := Counter{Value: 0}
    counter.Increment()
    fmt.Println("Counter Value:", counter.Value)
}
