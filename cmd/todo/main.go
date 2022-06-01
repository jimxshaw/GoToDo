package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/jimxshaw/GoToDo"
)

var listFileName = ".todo.json"

func main() {
	// Define and then parse the command line flags.
	// Return values are pointers.
	add := flag.Bool("add", false, "Add task to todo list")
	task := flag.String("task", "", "Task to be added to the list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item number to be marked as complete")

	flag.Parse()

	// Allow user to specify the file name.
	// E.g. export LIST_FILENAME=my-new-list.json
	if os.Getenv("LIST_FILENAME") != "" {
		listFileName = os.Getenv("LIST_FILENAME")
	}

	todoList := &todo.List{}

	if err := todoList.Get(listFileName); err != nil {
		// Standard Error is preferred over Standard Out for
		// errors as the user can filter them out more easily.
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		// List the items in list.
		// List type implemented Stringer interface.
		fmt.Print(todoList)
	// Item numbers start with 1.
	case *complete > 0:
		// Complete the item number then save the list.
		if err := todoList.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := todoList.Save(listFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		todoList.Add(*task)

		if err := todoList.Save(listFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		// Any non-flag argument(s) will be added as a new task.
		tasks, err := getTask(os.Stdin, flag.Args()...)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		todoList.Add(tasks)

		if err := todoList.Save(listFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// Invalid flags.
		fmt.Fprintln(os.Stderr, "Invalid option flag")
		os.Exit(1)
	}
}

// Figures out where to get the new task,
// either through arguments or STDIN.
func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	if len(scanner.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}

	return scanner.Text(), nil
}
