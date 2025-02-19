package main

import "fmt"

// Define a Person struct
type Person struct {
    Name string
    Age  int
}

// Method to increase person's age
func (p *Person) Birthday() {
    p.Age++
}

// Define a Car struct
type Car struct {
    Brand string
    Year  int
}

func main() {
    // Create a Person and call Birthday()
    person := Person{Name: "John", Age: 30}
    fmt.Println("Before Birthday:", person)
    person.Birthday()
    fmt.Println("After Birthday:", person)

    // Create a Car
    car := Car{Brand: "Toyota", Year: 2020}
    fmt.Println("Car:", car)
}
