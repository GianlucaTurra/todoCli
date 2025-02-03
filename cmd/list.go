package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

type row struct {
	id        int
	name      string
	completed int
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("No arguments expected")
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

		rows, err := tx.Query("select * from todos")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()
		table := table.New("ID", "Name", "Completed")
		table.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		for rows.Next() {
			var row row
			err = rows.Scan(&row.id, &row.name, &row.completed)
			if err != nil {
				log.Fatal(err)
			}
			// fmt.Println(row.id, row.name, row.completed == 1)
			table.AddRow(row.id, row.name, row.completed == 1)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
