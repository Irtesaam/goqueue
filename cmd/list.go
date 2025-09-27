package cmd

import (
	"fmt"
	"os"

	"github.com/Irtesaam/goqueue/internal/storage"
	"github.com/Irtesaam/goqueue/internal/todo"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todo items",
	Long: `Display all todo items in a beautiful table format.

This command shows:
- Item index (for reference in other commands)
- Task description
- Completion status (✅ completed, ❌ pending)
- Creation date
- Completion date (if completed)`,
	Aliases: []string{"ls", "show"},
	Run: func(cmd *cobra.Command, args []string) {
		// Load todos
		todos := todo.Todos{}
		store := storage.New[todo.Todos](GetTodoFile())

		if err := store.Load(&todos); err != nil {
			// If file doesn't exist, just show empty list
			if !os.IsNotExist(err) {
				fmt.Printf("Error loading todos: %v\n", err)
				os.Exit(1)
			}
		}

		if len(todos) == 0 {
			fmt.Println("📝 No todos found. Add some with 'goq add \"your task\"'")
			return
		}

		todos.Print()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
