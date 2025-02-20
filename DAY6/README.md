# Day 6: Pointers & Interfaces

## Learn:
- Pointers (`*` and `&`).
- Interfaces (Defining behaviors).

## Resources:
- [Pointers in Go](https://tour.golang.org/moretypes/1)
- [Interfaces in Go](https://gobyexample.com/interfaces)

## Exercises:
1. **Write a function using pointers to swap two numbers.**
2. **Create an interface `Shape` with an `Area()` method and implement it for `Circle` and `Rectangle`.**
3. **Use a pointer receiver to modify struct values inside a method.**

## Notes:
- Pointers allow modifying variables via memory addresses.
  - `*`: Get value
  - `&`: Get address
```
  package main
  import "fmt"

  type Person struct{
    Name string
    Age int
  }

  func changePerson(p *Person, nName string, nAge int) {
    p.Name = nName
    p.Age = nAge
  }

  func main() {
    var intern = &Person{"Kowala", 18}
    fmt.Println("Before change:", *intern) // Before change: {Kowala 18}
    changePerson(intern, "Haren Lin", 22)
    fmt.Println("After change:", *intern)  // After change: {Haren Lin 22}
  }
  ```
- Interfaces define behavior, making code more flexible. (like Polymorphism in C++)
  ``` 
  package main
  import "fmt"

  // define an interface of message sender
  type MessageSender interface {
    Send(content string)
  }

  // email sender part
  type EmailSender struct {
    Address string
  }

  func (s *EmailSender) Send(content string) {
    fmt.Println("Send:", content, "to", s.Address)
  }
  // email sender part
  // mobile sender part
  type SmsSender struct {
    Mobile string
  }

  func (s *SmsSender) Send(content string) {
    fmt.Println("Send:", content, "to", s.Mobile)
  }
  // mobile sender part
  func main() {
    var sender MessageSender // Interface Declaration
    sender = &EmailSender{Address: "test@gmail.com"}
    sender.Send("Hello, Email!") // Send: Hello, Email! to test@gmail.com
    sender = &SmsSender{Mobile: "0912345678"}
    sender.Send("Hello, Sms!") // Send: Hello, Sms! to 0912345678
  }
  ```
- Methods with pointer receivers modify struct values.
  ```
  package main
  import "fmt"

  type Hero struct {
    name   string
    health int
    power  int
  }

  func (h1 *Hero) Attack (h2 *Hero) {
    h2.health = h2.health - h1.power
  }

  func main() {
    var tony = &Hero{"IronMan", 200, 50}
    fmt.Println("Create Hero1:", *tony) // Create Hero1: {IronMan 200 50}
    var stephen = &Hero{"Dr.Strange", 100, 25}
    fmt.Println("Create Hero2:", *stephen) // Create Hero2: {Dr.Strange 100 25}

    tony.Attack(stephen)
    fmt.Println("After attack:", *tony, *stephen)
    // After attack: {IronMan 200 50} {Dr.Strange 50 25}
  }
  ```


