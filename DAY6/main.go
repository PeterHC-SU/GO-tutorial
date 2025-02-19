package main

import (
	"errors"
	"fmt"
	"math"
)

func validateNumber(value float64) error {
	if value <= 0 {
		return errors.New("數值必須大於 0")
	}
	return nil
}

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
func (c Circle) Area() (float64, error) {
	if err := validateNumber(c.Radius); err != nil {
		return 0, err
	}
	return math.Pi * c.Radius * c.Radius, nil
}

// Implement Area() for Rectangle
func (r Rectangle) Area() (float64, error) {
	if err := validateNumber(r.Width); err != nil {
		return 0, err
	}
	if err := validateNumber(r.Height); err != nil {
		return 0, err
	}
	return r.Width * r.Height, nil
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

	if area, err := c.Area(); err != nil {
		fmt.Println("Circle Area Error:", err)
	} else {
		fmt.Println("Circle Area:", area)
	}

	if area, err := r.Area(); err != nil {
		fmt.Println("Rectangle Area Error:", err)
	} else {
		fmt.Println("Rectangle Area:", area)
	}

	// Pointer receiver method
	counter := Counter{Value: 0}
	counter.Increment()
	fmt.Println("Counter Value:", counter.Value)
}
