package todo

import (
	"fmt"
	"time"
)

// This struct represents a ToDo item.
type item struct {
	Task           string
	Done           bool
	CreationDate   time.Time
	CompletionDate time.Time
}

// This is a list of ToDo items.
type List []item

func (l *List) Add(task string) {
	newItem := item{
		Task:           task,
		Done:           false,
		CreationDate:   time.Now(),
		CompletionDate: time.Time{},
	}

	*l = append(*l, newItem)
}

func (l *List) Delete(item int) error {
	list := *l

	if item < 0 || item > len(list) {
		return fmt.Errorf("Item %d does not exist", item)
	}

	*l = append(list[:item-1], list[item:]...)

	return nil
}

func (l *List) Complete(item int) error {
	list := *l

	if item <= 0 || item > len(list) {
		return fmt.Errorf("Item %d does not exist", item)
	}

	// Arrays are 0 based and need to be adjusted.
	list[item-1].Done = true
	list[item-1].CompletionDate = time.Now()

	return nil
}
