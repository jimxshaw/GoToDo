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
