package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building application...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "test task 1"

	directory, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(directory, binName)

	t.Run("AddNewTaskFromArguments", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	task2 := "Testing another task"

	t.Run("AddNewTaskFromSTDIN", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")

		cmdStdIn, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}

		io.WriteString(cmdStdIn, task2)
		cmdStdIn.Close()

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")

		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal()
		}

		expected := fmt.Sprintf(" 1: %s\n 2: %s\n", task, task2)

		if expected != string(output) {
			t.Errorf("Expected %q, got %q", expected, output)
		}
	})

	t.Run("CompleteTask", func(t *testing.T) {
		// Mark task as complete.
		complete := exec.Command(cmdPath, "-complete", "1")

		if err := complete.Run(); err != nil {
			t.Fatal(err)
		}

		// List out the tasks.
		list := exec.Command(cmdPath, "-list")

		output, err := list.CombinedOutput()
		if err != nil {
			t.Fatal()
		}

		expected := fmt.Sprintf("X 1: %s\n", task)

		// If there's only 1 task and it is marked as complete then
		// the output of the list command would be blank.
		if expected != string(output) {
			t.Errorf("Expected %q, got %q", expected, output)
		}

	})
}
