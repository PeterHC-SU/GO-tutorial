package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

// Struct for JSON encoding/decoding
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    // Writing to a file
    content := "Hello, Go file handling!"
    err := ioutil.WriteFile("example.txt", []byte(content), 0644)
    if err != nil {
        fmt.Println("Error writing file:", err)
        return
    }

    // Reading from a file
    data, err := ioutil.ReadFile("example.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    fmt.Println("File Content:", string(data))

    // Writing JSON to a file
    person := Person{Name: "Alice", Age: 25}
    jsonData, err := json.MarshalIndent(person, "", "  ")
    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return
    }
    ioutil.WriteFile("person.json", jsonData, 0644)

    // Reading JSON from a file
    jsonFile, err := os.ReadFile("person.json")
    if err != nil {
        fmt.Println("Error reading JSON file:", err)
        return
    }
    var decodedPerson Person
    json.Unmarshal(jsonFile, &decodedPerson)
    fmt.Println("Decoded JSON:", decodedPerson)
}
