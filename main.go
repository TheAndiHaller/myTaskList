package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Task struct {
	Done      bool
	Text      string
	Project   string
	Type      string
	Added     time.Time
	Completed time.Time
}

func main() {
	tasks := readTasks("list.md")

	printTasks(tasks)

	saveTasks("list.md", tasks)
}

// Print parsed tasks
func printTasks(tasks []Task) {
	for i, t := range tasks {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, t.Text)
	}
}

// Function to read the file and load into memory (slice)
func readTasks(filename string) []Task {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}

	defer file.Close()

	var tasks []Task

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "- [ ]") {
			tasks = append(tasks, Task{Done: false, Text: strings.TrimSpace(line[5:])})
		} else if strings.HasPrefix(line, "- [x]") {
			tasks = append(tasks, Task{Done: true, Text: strings.TrimSpace(line[5:])})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	return tasks
}

func saveTasks(filename string, tasks []Task) error {
	// Create (or overwrite) the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Optional: add a title
	_, err = file.WriteString("# ToDo\n")
	if err != nil {
		return err
	}

	// Write each task
	for _, t := range tasks {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		line := fmt.Sprintf("- %s %s\n", status, t.Text)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}

	return nil
}
