package cmd

import (
	"fmt"
	"os"

	"github.com/Irtesaam/goqueue/internal/storage"
	"github.com/Irtesaam/goqueue/internal/todo"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new todo item",
	Long: `Add a new todo item to your list.

Examples:
  goq add "Buy groceries"
  goq add "Complete the project report"
  goq add "Call mom"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Join all arguments to handle multi-word tasks
		title := ""
		for i, arg := range args {
			if i > 0 {
				title += " "
			}
			title += arg
		}

		// Load todos
		todos := todo.Todos{}
		store := storage.New[todo.Todos](GetTodoFile())
		store.LoadOrInitialize(&todos)

		// Add new todo
		todos.Add(title)

		// Save todos
		if err := store.Save(todos); err != nil {
			fmt.Printf("Error saving todo: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ Added: \"%s\"\n", title)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
