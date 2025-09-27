package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Irtesaam/goqueue/internal/storage"
	"github.com/Irtesaam/goqueue/internal/todo"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete [index]",
	Short:   "Delete a todo item",
	Aliases: []string{"del", "rm", "remove"},
	Long: `Delete a todo item by its index.

Use 'goq list' to see the index numbers of your tasks.

Examples:
  goq delete 0    # Delete the first task
  goq del 2       # Delete the third task (using alias)`,
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

		if err := store.LoadOrInitialize(&todos); err != nil {
			fmt.Printf("Error loading todos: %v\n", err)
			os.Exit(1)
		}

		if len(todos) == 0 {
			fmt.Println("📝 No todos found to delete.")
			return
		}

		// Store the title before deletion for confirmation message
		var title string
		if index >= 0 && index < len(todos) {
			title = todos[index].Title
		}

		// Delete todo
		if err := todos.Delete(index); err != nil {
			fmt.Printf("❌ %v\n", err)
			os.Exit(1)
		}

		// Save todos
		if err := store.Save(todos); err != nil {
			fmt.Printf("Error saving todos: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("🗑️  Deleted: \"%s\"\n", title)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
