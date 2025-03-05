package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var updateAll bool
var statement string

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark task as completed",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 && !updateAll {
			log.Fatal("You need to choose which task should be completed")
		}

		db, err := sql.Open("sqlite3", "./todo.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		if updateAll {
			statement = "update todos set completed = 1"
		} else {
			statement = fmt.Sprintf("update todos set completed = 1 where id = %s", args[0])
		}

		_, err = db.Exec(statement)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	completeCmd.Flags().BoolVarP(&updateAll, "all", "a", false, "Mark all tasks as completed")
}
