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
var inputTask string

func main() {
	flag.StringVar(&inputTask, "t", "Clean your room", "Add a task to complete.")
	flag.Parse()
	fmt.Println("New task added: " + inputTask)

	f, err := os.OpenFile(taskListFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	check(err)
	defer f.Close()

	newTask := createTask(inputTask, taskListFileName)
	write(f, newTask)
}

func createTask(inputTask string, fileName string) Task {
	return Task{
		Order:  getLinesQty(fileName) + 1,
		Name:   inputTask,
		IsDone: false,
	}
}

func getLinesQty(fileName string) int {
	file, err := os.Open(fileName)

	check(err)

	fileScanner := bufio.NewScanner(file)
	lineCount := 0

	for fileScanner.Scan() {
		fmt.Println("incrementing line count.")
		lineCount++
	}
	file.Close()
	fmt.Printf("Number of lines: %d\n", lineCount)
	return lineCount
}

func write(f *os.File, task Task) {
	line := createFileLine(task)
	w := bufio.NewWriter(f)
	bw, err := w.WriteString(line)
	check(err)
	fmt.Printf("wrote %d bytes.", bw)
	w.Flush()

}

func createFileLine(task Task) string {
	return fmt.Sprintf("%d,%s,%t\n", task.Order, task.Name, task.IsDone)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
