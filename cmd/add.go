package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var dueDate string
var description string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new TODO",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("This command is supposed to receive one argument: the task's name")
			return
		}

		db, err := sql.Open("sqlite3", "./todo.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		statement, err := tx.Prepare("insert into todos (name, description, dueDate) values (?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer statement.Close()

		_, err = statement.Exec(args[0], description, dateToUnix(dueDate))
		if err != nil {
			log.Fatal(err)
		}

		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVar(&dueDate, "due-date", "", "Set a deadline for yourself")
	addCmd.Flags().StringVarP(&description, "description", "d", "", "A verbose description for the task")
}

func dateToUnix(date string) int64 {

	if date == "" {
		return time.Now().Unix()
	}

	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatal(err)
	}
	return t.Unix()
}
