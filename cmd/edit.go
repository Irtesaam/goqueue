package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Irtesaam/goqueue/internal/storage"
	"github.com/Irtesaam/goqueue/internal/todo"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [index] [new description]",
	Short: "Edit a todo item",
	Long: `Edit the description of a todo item by its index.

Use 'goq list' to see the index numbers of your tasks.

Examples:
  goq edit 0 "Updated task description"
  goq edit 2 "Buy milk instead of bread"`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("❌ Invalid index: %s. Please provide a valid number.\n", args[0])
			os.Exit(1)
		}

		// Join all arguments after the index to handle multi-word descriptions
		newTitle := ""
		for i, arg := range args[1:] {
			if i > 0 {
				newTitle += " "
			}
			newTitle += arg
		}

		// Load todos
		todos := todo.Todos{}
		store := storage.New[todo.Todos](GetTodoFile())

		if err := store.LoadOrInitialize(&todos); err != nil {
			fmt.Printf("Error loading todos: %v\n", err)
			os.Exit(1)
		}

		if len(todos) == 0 {
			fmt.Println("📝 No todos found to edit.")
			return
		}

		// Store the old title for confirmation message
		var oldTitle string
		if index >= 0 && index < len(todos) {
			oldTitle = todos[index].Title
		}

		// Edit todo
		if err := todos.Edit(index, newTitle); err != nil {
			fmt.Printf("❌ %v\n", err)
			os.Exit(1)
		}

		// Save todos
		if err := store.Save(todos); err != nil {
			fmt.Printf("Error saving todos: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✏️  Updated: \"%s\" → \"%s\"\n", oldTitle, newTitle)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
