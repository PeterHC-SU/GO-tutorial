# Day 8: File Handling & JSON Processing

## Learn:
- Read/Write files in Go.
- JSON encoding/decoding.

## Resources:
- [Reading and Writing Files](https://gobyexample.com/reading-files)
- [JSON in Go](https://gobyexample.com/json)

## Exercises:
1. **Write a program that reads a file and prints its content.**
2. **Create a Go struct and serialize it to JSON.**
3. **Read JSON data from a file and parse it.**

## Notes:
| Function              | Purpose                                | Example Use Case               |
|----------------------|--------------------------------------|--------------------------------|
| `ioutil.WriteFile()`  | Writes data to a file.              | Save JSON data to a file.      |
| `ioutil.ReadFile()`   | Reads the full file content.        | Load JSON data from a file.    |
| `json.Marshal()`      | Converts Go struct to compact JSON. | Serialize data before storing. |
| `json.MarshalIndent()` | Converts Go struct to **formatted** JSON. | Pretty-print JSON for readability. |
| `json.Unmarshal()`    | Parses JSON back into Go struct.    | Deserialize data after loading. |

- `ioutil.WriteFile()` and `ioutil.ReadFile()` handle file I/O.
```go 
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Define a struct
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	user := User{Username: "Alice", Email: "alice@example.com"}

	// Convert struct to JSON
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// Write JSON to file
	err = ioutil.WriteFile("user.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data written to file.")

	// Read JSON from file
	data, err := ioutil.ReadFile("user.json")
	if err != nil {
		log.Fatal(err)
	}

	// Convert JSON back to struct
	var newUser User
	err = json.Unmarshal(data, &newUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read from file: %+v\n", newUser)
}
```
- `json.Marshal()` encodes Go structs to JSON.
```go
person := Person{Name: "John Doe", Age: 30, Address: "123 Main St"}
jsonData, err := json.Marshal(person)
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(jsonData)) // OUTPUT: {"name":"John Doe","age":30,"address":"123 Main St"}
```
- `json.Unmarshal()` decodes JSON into Go structs.
```go
var newPerson Person
err = json.Unmarshal(jsonData, &newPerson)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("%+v\n", newPerson) // OUTPUT: {Name:John Doe Age:30 Address:123 Main St}
```
- `json.MarshalIndent()` works like `json.Marshal()`, but formats the output with indentation.
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Define a struct
type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	product := Product{Name: "Laptop", Price: 1299.99}

	// Pretty-print JSON
	jsonData, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
	/* OUTPUT:
	{
	  "name": "Laptop",
	  "price": 1299.99
	}
	*/
}
```