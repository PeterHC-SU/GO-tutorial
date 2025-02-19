package main

import (
	"math"
	"testing"
)

// TEST Swap
func TestSwap(t *testing.T) {
	a, b := -10, 20
	swap(&a, &b)
	if a != 20 || b != -10 {
		t.Errorf("swap() failed: expected (20, -10), got (%d, %d)", a, b)

	}
}

// TEST Circle Area()
func TestCircleArea(t *testing.T) {
	circle := Circle{Radius: 5}
	expected := math.Pi * 5 * 5
	got, err := circle.Area()
	if err != nil {
		t.Errorf("Circle Area() failed: unexpected error %v", err)
	}

	if got != expected {
		t.Errorf("Circle Area() failed: expected %.2f, got %.2f", expected, got)
	}
}

// TEST Rectangle Area()
func TestRectangleArea(t *testing.T) {
	rectangle := Rectangle{Width: 4, Height: 6}
	expected := 4 * 6
	got, err := rectangle.Area()
	if err != nil {
		t.Errorf("Rectangle Area() failed: unexpected error %v", err)
	}
	if got != float64(expected) {
		t.Errorf("Rectangle Area() failed: expected %.2f, got %.2f", float64(expected), got)
	}
}

// TEST Counter Increment()
func TestCounterIncrement(t *testing.T) {
	init := 0
	counter := Counter{Value: init}
	counter.Increment()

	if counter.Value != init+1 {
		t.Errorf("Increment() failed: expected 1, got %d", counter.Value)
	}
}
