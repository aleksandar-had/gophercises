package cmd

import (
	"fmt"
	"strconv"

	"github.com/aleksandar-had/gophercises/task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Completes a given task.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Failed to parse %s arg.\n", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Error while fetching tasks:", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Couldn't mark \"%d\" for completion. Error: %s\n", id, err)
			}
			fmt.Printf("Marked \"%d\" as completed.\n", id)
		}

		fmt.Println("Tasks", ids, "marked as done.")
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
