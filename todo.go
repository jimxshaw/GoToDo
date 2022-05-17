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
