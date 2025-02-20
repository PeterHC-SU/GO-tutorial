package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Task represents a todo item
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Completed   bool      `json:"completed"`
	Deadline    time.Time `json:"deadline"`
	Priority    int       `json:"priority"` // 1: High, 2: Medium, 3: Low
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
}

// TaskList represents the list of tasks
type TaskList struct {
	Tasks []Task `json:"tasks"`
}

const filename = "tasks.json"

func main() {
	// Define command line flags
	add := flag.String("add", "", "Add a new task")
	complete := flag.Int("complete", 0, "Mark a task as completed") // complete task {id}
	del := flag.Int("del", 0, "Delete a task")                      // delete task {id}
	list := flag.String("list", "all", "List tasks (all/pending/completed)")
	deadline := flag.String("deadline", "", "Set deadline for new task (YYYY-MM-DD)")
	priority := flag.Int("priority", 2, "Set priority for new task (1: High, 2: Medium, 3: Low)")

	flag.Parse()

	// Load existing tasks
	tasks := loadTasks()

	// Handle different commands
	switch {
	case *add != "":
		addTask(tasks, *add, *deadline, *priority)
	case *complete > 0:
		completeTask(tasks, *complete)
	case *del > 0:
		deleteTask(tasks, *del)
	default:
		listTasks(tasks, *list)
	}
}

// loadTasks loads tasks from JSON file
func loadTasks() *TaskList {
	tasks := &TaskList{}

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return tasks
	}

	// Read file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return tasks
	}

	// Parse JSON
	err = json.Unmarshal(data, tasks)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return tasks
	}

	return tasks
}

// saveTasks saves tasks to JSON file
func saveTasks(tasks *TaskList) {
	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		return
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
	}
}

// addTask adds a new task
func addTask(tasks *TaskList, title string, deadlineStr string, priority int) {
	var deadline time.Time
	var err error

	if deadlineStr != "" {
		deadline, err = time.Parse("2006-01-02", deadlineStr)
		if err != nil {
			fmt.Println("Invalid deadline format. Use YYYY-MM-DD")
			return
		}
	}

	newTask := Task{
		ID:        len(tasks.Tasks) + 1,
		Title:     title,
		Priority:  priority,
		CreatedAt: time.Now(),
	}

	if !deadline.IsZero() {
		newTask.Deadline = deadline
	}

	tasks.Tasks = append(tasks.Tasks, newTask)
	saveTasks(tasks)
	fmt.Printf("Added task: %s\n", title)
}

// completeTask marks a task as completed
func completeTask(tasks *TaskList, id int) {
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == id {
			tasks.Tasks[i].Completed = true
			tasks.Tasks[i].CompletedAt = time.Now()
			saveTasks(tasks)
			fmt.Printf("Marked task %d as completed\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", id)
}

// deleteTask deletes a task
func deleteTask(tasks *TaskList, id int) {
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == id {
			tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
			saveTasks(tasks)
			fmt.Printf("Deleted task %d\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", id)
}

// listTasks displays tasks based on their status
func listTasks(tasks *TaskList, filter string) {
	if len(tasks.Tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Println("\nTasks:")
	fmt.Println("----------------------------------------")

	for _, task := range tasks.Tasks {
		// Apply filter
		if filter == "pending" && task.Completed {
			continue
		}
		if filter == "completed" && !task.Completed {
			continue
		}

		// Get status symbol
		status := "[ ]"
		if task.Completed {
			status = "[✓]"
		}

		// Get priority string // change format int to string
		priority := "Medium"
		switch task.Priority {
		case 1:
			priority = "High"
		case 3:
			priority = "Low"
		}

		// Format deadline
		deadline := "No deadline"
		if !task.Deadline.IsZero() {
			deadline = task.Deadline.Format("2025-01-02")
		}

		fmt.Printf("%s #%d: %s\n", status, task.ID, task.Title)
		fmt.Printf("   Priority: %s, Deadline: %s\n", priority, deadline)

		if task.Completed {
			fmt.Printf("   Completed at: %s\n", task.CompletedAt.Format("2025-01-02 12:03:04"))
		}
		fmt.Println("----------------------------------------")
	}
}
