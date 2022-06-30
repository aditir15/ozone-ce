/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

type project struct {
	id          int
	name        string
	description string
}

// createProjectCmd represents the createProject command
var createProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createProject called with arguments", args)
		fmt.Println("the project name is : ", args[0])

		//1. create a database

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
		newProject := project{
			name:        args[0],
			description: args[1],
		}

		addProject(db, newProject)
	},
}

func init() {
	createCmd.AddCommand(createProjectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createProjectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createProjectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addProject(db *sql.DB, newProject project) {

	stmt, err := db.Prepare("INSERT INTO projects (id, name, description) VALUES (?, ?, ?)")
	checkErr(err)
	result, err := stmt.Exec(nil, newProject.name, newProject.description)
	checkErr(err)
	fmt.Printf("result %v", result)
	defer stmt.Close()

	fmt.Printf("Added %v %v \n", newProject.name, newProject.description)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
