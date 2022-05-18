package todo

import (
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
