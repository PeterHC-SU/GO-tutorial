package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Task struct {
    Name      string `json:"name"`
    Completed bool   `json:"completed"`
}

func main() {
    tasks := []Task{
        {Name: "Learn Go", Completed: false},
        {Name: "Build a CLI app", Completed: false},
    }

    file, _ := json.MarshalIndent(tasks, "", "  ")
    os.WriteFile("tasks.json", file, 0644)

    fmt.Println("Tasks saved to tasks.json")
}
