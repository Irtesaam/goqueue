package cmd

import (
	"fmt"
	"os"
	"strconv"

	"todo/internal/storage"
	"todo/internal/todo"

	"github.com/spf13/cobra"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle [index]",
	Short: "Toggle completion status of a todo item",
	Long: `Toggle the completion status of a todo item by its index.
This will mark a pending task as completed, or a completed task as pending.

Use 'goq list' to see the index numbers of your tasks.

Examples:
  goq toggle 0    # Toggle the first task
  goq toggle 2    # Toggle the third task`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("❌ Invalid index: %s. Please provide a valid number.\n", args[0])
			os.Exit(1)
		}

		// Load todos
		todos := todo.Todos{}
		store := storage.New[todo.Todos](GetTodoFile())

		if err := store.Load(&todos); err != nil {
			fmt.Printf("Error loading todos: %v\n", err)
			os.Exit(1)
		}

		if len(todos) == 0 {
			fmt.Println("📝 No todos found to toggle.")
			return
		}

		// Store the title and current status for confirmation message
		var title string
		var wasCompleted bool
		if index >= 0 && index < len(todos) {
			title = todos[index].Title
			wasCompleted = todos[index].Completed
		}

		// Toggle todo
		if err := todos.Toggle(index); err != nil {
			fmt.Printf("❌ %v\n", err)
			os.Exit(1)
		}

		// Save todos
		if err := store.Save(todos); err != nil {
			fmt.Printf("Error saving todos: %v\n", err)
			os.Exit(1)
		}

		status := "completed ✅"
		if wasCompleted {
			status = "pending ❌"
		}
		fmt.Printf("🔄 Toggled \"%s\" to %s\n", title, status)
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}
