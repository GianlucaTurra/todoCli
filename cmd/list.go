package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

type row struct {
	id          int
	name        string
	description string
	dueDate     int64
	completed   int
}

var completed bool
var uncompleted bool
var listStatement string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			fmt.Println("No arguments expected")
			return
		}

		if completed && uncompleted {
			log.Fatal("Choose completed or not... Default is actually all")
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

		if allCompleted {
			statement = "select * from todos where completed = 1"
		} else if uncompleted {
			statement = "select * from todos where completed = 0"
		} else {
			statement = "select * from todos"
		}

		rows, err := tx.Query(statement)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		headerFmt := color.New(color.FgGreen).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()
		table := table.New("ID", "Name", "Description", "Due Date", "Completed")
		table.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		for rows.Next() {
			var row row
			err = rows.Scan(&row.id, &row.name, &row.description, &row.dueDate, &row.completed)
			if err != nil {
				log.Fatal(err)
			}
			table.AddRow(row.id, row.name, row.description, time.Unix(row.dueDate, 0).Format("2006-01-02"), row.completed == 1)
		}
		table.Print()
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&completed, "completed", "c", false, "List only completed tasks")
	listCmd.Flags().BoolVarP(&uncompleted, "uncompleted", "u", false, "List only uncompleted tasks")
}
