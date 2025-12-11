package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/example/hello/internal/todo"
)

func usage() {
	prog := filepath.Base(os.Args[0])
	fmt.Printf("Usage:\n")
	fmt.Printf("  %s add <task text>       Add a new task\n", prog)
	fmt.Printf("  %s list [--all]         List tasks (default: pending only)\n", prog)
	fmt.Printf("  %s done <id>            Mark task <id> done\n", prog)
	fmt.Printf("  %s rm <id>              Remove task <id>\n", prog)
	fmt.Printf("  %s help                 Show this help\n", prog)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "add: missing task text")
			os.Exit(1)
		}
		text := strings.Join(os.Args[2:], " ")
		t, err := todo.AddTask("", text)
		if err != nil {
			fmt.Fprintln(os.Stderr, "add error:", err)
			os.Exit(1)
		}
		fmt.Printf("added %d: %s\n", t.ID, t.Text)

	case "list":
		fs := flag.NewFlagSet("list", flag.ExitOnError)
		all := fs.Bool("all", false, "show all tasks, including done ones")
		_ = fs.Parse(os.Args[2:])
		tasks, err := todo.ListTasks("")
		if err != nil {
			fmt.Fprintln(os.Stderr, "list error:", err)
			os.Exit(1)
		}
		for _, t := range tasks {
			if !*all && t.Done {
				continue
			}
			mark := ' '
			if t.Done {
				mark = 'x'
			}
			fmt.Printf("[%c] %d: %s\n", mark, t.ID, t.Text)
		}

	case "done":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "done: missing id")
			os.Exit(1)
		}
		id, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "done: invalid id")
			os.Exit(1)
		}
		if err := todo.MarkDone("", id); err != nil {
			fmt.Fprintln(os.Stderr, "done error:", err)
			os.Exit(1)
		}
		fmt.Printf("marked %d done\n", id)

	case "rm":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "rm: missing id")
			os.Exit(1)
		}
		id, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "rm: invalid id")
			os.Exit(1)
		}
		if err := todo.RemoveTask("", id); err != nil {
			fmt.Fprintln(os.Stderr, "rm error:", err)
			os.Exit(1)
		}
		fmt.Printf("removed %d\n", id)

	case "help":
		usage()

	default:
		fmt.Fprintln(os.Stderr, "unknown command:", cmd)
		usage()
		os.Exit(1)
	}
}
