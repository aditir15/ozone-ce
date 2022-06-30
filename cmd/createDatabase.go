/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"
)

// createDatabaseCmd represents the createDatabase command
var createDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createDatabase called")
		// Connect to database
		db, err := sql.Open("sqlite3", "./names.db")
		checkErr(err)
		// defer close
		defer db.Close()

		//2. create a table
		stmt, err := db.Prepare(`CREATE TABLE "projects" (
			"id" INTEGER,
			"name"	TEXT,
			"description"	TEXT,
			PRIMARY KEY("id" AUTOINCREMENT)
		)`)
		checkErr(err)
		stmt.Exec()
		defer stmt.Close()
	},
}

func init() {
	createCmd.AddCommand(createDatabaseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createDatabaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createDatabaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
