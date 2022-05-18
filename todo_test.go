package todo_test

import (
	"testing"

	todo "github.com/jimxshaw/GoToDo"
)

func TestAdd(t *testing.T) {
	list := todo.List{}

	task := "Clean House"
	list.Add(task)

	if list[0].Task != task {
		t.Errorf("Expected %q, got %q", task, list[0].Task)
	}
}

func TestComplete(t *testing.T) {
	list := todo.List{}

	task := "Clean House"
	list.Add(task)

	if list[0].Task != task {
		t.Errorf("Expected %q, got %q", task, list[0].Task)
	}

	if list[0].Done {
		t.Errorf("New task should not be done but it is.")
	}

	list.Complete(1)

	if !list[0].Done {
		t.Errorf("New task should be done but it is not.")
	}
}
