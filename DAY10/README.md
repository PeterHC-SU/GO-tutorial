# Day 10: Mini Project - To-Do CLI App (Optional)

## Project:
- Apply everything learned to a small project.

### **Project Idea: 📝 To-Do CLI App**
- Add/Delete tasks.
- Store tasks in a file.
- Display pending/completed tasks.

## Resources:
- [Build a CLI App in Go](https://freshman.tech/todo-list-cli/)

## Exercises:
1. **Modify the app to save tasks as JSON.**
2. **Add a command-line flag for filtering completed tasks.**
3. **Extend the app with additional features (e.g., deadlines).**

## Notes:
- Use JSON for persistent storage.
- Parse command-line arguments for more functionality.
- Expand the app to track deadlines and priorities.

## Start:
```
# build
go build -o todo

# Add priority high task
./todo -add "project report" -priority 1 -deadline "2025-03-01"
./todo -add "Golang tutorial" -priority 1 -deadline "2025-02-28"

# Add priority medium task
./todo -add "weekly meeting" -priority 2 -deadline "2025-02-25"
./todo -add "document" -priority 2 -deadline "2025-03-05"

# Add priority low task
./todo -add "side project" -priority 3 -deadline "2025-03-15"

# Tag completed
./todo -complete 2
./todo -complete 4

# list all tasks (all/ priority/ pending/completed)
./todo -list all
```
