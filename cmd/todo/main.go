package main

import (
	"fmt"
	"os"
	"strings"

	todo "github.com/jimxshaw/GoToDo"
)

const listFileName = ".todo.json"

func main() {
	list := &todo.List{}

	if err := list.Get(listFileName); err != nil {
		// Standard Error is preferred over Standard Out for
		// errors as the user can filter them out more easily.
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// User arguments will determine what needs to be done next.
	switch {
	// Print the list if there aren't more than 1 argument.
	case len(os.Args) == 1:
		for _, item := range *list {
			fmt.Println(item.Task)
		}
	default:
		// The first argument is always the program name.
		task := strings.Join(os.Args[1:], " ")

		list.Add(task)

		if err := list.Save(listFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
