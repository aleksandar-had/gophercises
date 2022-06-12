package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aleksandar-had/gophercises/task/cmd"
	"github.com/aleksandar-had/gophercises/task/db"
)

func main() {
	home, _ := os.UserHomeDir()

	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	cmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
