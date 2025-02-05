package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var all bool
var allCompleted bool
var deleteStatement string

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a record",
	Long:  `Delete a record from the database based on the task name`,
	Run: func(cmd *cobra.Command, args []string) {

		if all && allCompleted {
			log.Fatal("Delete all tasks and all completed tasks... which one?")
		}

		if len(args) != 1 && (all || allCompleted) {
			log.Fatal("Select a tasks of user `all`/`allCompleted`")
		}

		db, err := sql.Open("sqlite3", "./todo.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		if all {
			deleteStatement = "delete from todos"
		} else if allCompleted {
			deleteStatement = "delete from todos where completed = 1"
		} else {
			deleteStatement = fmt.Sprintf("delete from todos where name = '%s'", args[0])
		}

		_, err = db.Exec(deleteStatement)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().BoolVarP(&all, "all", "a", false, "Delete all tasks")
	deleteCmd.Flags().BoolVarP(&allCompleted, "all-comp", "c", false, "Delete all completed tasks")
}
