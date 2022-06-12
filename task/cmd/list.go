package cmd

import (
	"fmt"
	"os"

	"github.com/aleksandar-had/gophercises/task/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Tasks couldn't be fetched!")
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks to complete! Chill is acceptable.")
			return
		}
		fmt.Println("You have the following tasks:")
		for id, task := range tasks {
			fmt.Printf("%d. %s\n", id+1, task.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
