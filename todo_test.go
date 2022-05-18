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

func TestDelete(t *testing.T) {
	list := todo.List{}

	tasks := []string{
		"Do Laundry",
		"Clean House",
		"Mow Lawn",
		"Fold Clothes",
		"Cook Dinner",
		"Wash dishes",
	}

	for _, task := range tasks {
		list.Add(task)
	}

	if list[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q", tasks[0], list[0].Task)
	}

	list.Delete(3)

	if len(list) != 5 {
		t.Errorf("Expected list length of %d but got %d", 5, len(list))
	}

	if list[2].Task != tasks[3] {
		t.Errorf("Expected %q but got %q", tasks[3], list[2].Task)
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
