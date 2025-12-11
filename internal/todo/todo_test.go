package todo

import (
	"path/filepath"
	"testing"
)

func TestAddListMarkRemove(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "todos.json")

	// Add two tasks
	t1, err := AddTask(path, "first task")
	if err != nil {
		t.Fatalf("AddTask 1 failed: %v", err)
	}
	t2, err := AddTask(path, "second task")
	if err != nil {
		t.Fatalf("AddTask 2 failed: %v", err)
	}

	tasks, err := ListTasks(path)
	if err != nil {
		t.Fatalf("ListTasks failed: %v", err)
	}
	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}

	// Mark first as done
	if err := MarkDone(path, t1.ID); err != nil {
		t.Fatalf("MarkDone failed: %v", err)
	}
	tasks, _ = ListTasks(path)
	var doneCount int
	for _, tt := range tasks {
		if tt.Done {
			doneCount++
		}
	}
	if doneCount != 1 {
		t.Fatalf("expected 1 done task, got %d", doneCount)
	}

	// Remove second
	if err := RemoveTask(path, t2.ID); err != nil {
		t.Fatalf("RemoveTask failed: %v", err)
	}
	tasks, _ = ListTasks(path)
	if len(tasks) != 1 {
		t.Fatalf("expected 1 task after remove, got %d", len(tasks))
	}
}
