package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/jimxshaw/GoToDo"
)

const listFileName = ".todo.json"

func main() {
	// Define and then parse the command line flags.
	// Return values are pointers.
	task := flag.String("task", "", "Task to be added to the list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item number to be marked as complete")

	flag.Parse()

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
	default:
		// Invalid flags.
		fmt.Fprintln(os.Stderr, "Invalid option flag")
		os.Exit(1)
	}
}
