package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/aleksandar-had/gophercises/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Task couldn't be added to the list! An error occured!")
			os.Exit(1)
		}
		fmt.Printf("Added \"%s\" to the task list.\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
