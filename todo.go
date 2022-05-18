package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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

// Get method opens the file, decodes the
// JSON and parses it into a list.
func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

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

	// Get all the items in the list up to but not
	// including the item to be deleted and then appending
	// all items after the item to be deleted.
	// Using ... to unpack the slice as the append argument cannot
	// take []item but rather only item.
	*l = append(list[:item-1], list[item:]...)

	return nil
}

// Save method encodes the list as JSON and saves it to file.
func (l *List) Save(filename string) error {
	json, err := json.Marshal(l)

	if err != nil {
		return err
	}

	// A FileMode represents a file's mode and permission bits.
	// The bits have the same definition on all systems,
	// so that information about files can be moved from one
	// system to another portably. Not all bits apply to all systems.
	// The only required bit is ModeDir for directories.
	// 0644 means (6) file's owner can read & write,
	// (4) users in the same group as the file's owner can read and
	// (4) all users can read.
	return os.WriteFile(filename, json, 0644)
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
