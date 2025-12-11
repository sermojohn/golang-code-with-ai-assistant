package todo

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// TodoFilePath returns the default file path for storing tasks.
// It honors the TODO_FILE environment variable if set.
func TodoFilePath() string {
	if v := os.Getenv("TODO_FILE"); v != "" {
		return v
	}

	if currentDif := os.Getenv("PWD"); currentDif != "" {
		return filepath.Join(currentDif, ".todos.json")
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return ".todo.json"
	}
	return filepath.Join(home, ".todo.json")
}

// loadTasks reads and unmarshals tasks from the given file path.
// If the file does not exist, an empty slice is returned.
func loadTasks(path string) ([]Task, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return []Task{}, nil
	}
	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// saveTasks marshals tasks and atomically writes them to the given path.
func saveTasks(path string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)
	tmp := filepath.Join(dir, ".todo.json.tmp")
	if err := os.WriteFile(tmp, data, 0o600); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
