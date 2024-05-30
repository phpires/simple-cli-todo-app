package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Task struct {
	Name   string
	Order  int
	IsDone bool
}

var taskListFileName string = "task-list"
var newTask string
var todoList Task

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createFileLine(task Task) string {
	return fmt.Sprintf("%d %s %t", task.Order, task.Name, task.IsDone)
}

func write(f *os.File, task Task) {
	line := createFileLine(task)
	w := bufio.NewWriter(f)
	bw, err := w.WriteString(line)
	check(err)
	fmt.Printf("wrote %d bytes.", bw)
	w.Flush()

}

func main() {
	flag.StringVar(&newTask, "t", "Clean your room", "Add a task to complete.")
	flag.Parse()

	todoList.Name = newTask
	todoList.Order = 1
	todoList.IsDone = false

	fmt.Println("New task added: " + newTask)

	f, err := os.Create(taskListFileName)
	check(err)

	defer f.Close()
	write(f, todoList)
}
