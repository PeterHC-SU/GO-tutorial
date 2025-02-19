package main

import (
	"testing"
)

// TEST Person Birthday()
func TestBirthday(t *testing.T) {
	person := Person{Name: "Alice", Age: 25}
	person.Birthday()

	if person.Age != 26 {
		t.Errorf("Birthday() failed: expected age 26, got %d", person.Age)
	}
}

// TEST Car struct
func TestCarInitialization(t *testing.T) {
	car := Car{Brand: "Honda", Year: 2022}

	if car.Brand != "Honda" || car.Year != 2022 {
		t.Errorf("Car initialization failed: expected (Honda, 2022), got (%s, %d)", car.Brand, car.Year)
	}
}
