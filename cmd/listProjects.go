/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// listProjectsCmd represents the listProjects command
var listProjectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listProjects called")
		// Connect to database
		db, err := sql.Open("sqlite3", "./names.db")
		checkErr(err)
		// defer close
		defer db.Close()

		//read a record
		fmt.Println("Reading Records.......... ")
		project := searchForProject(db)
		for _, ourProject := range project {
			fmt.Printf("\n----\nID: %d\nName: %s\nDescription: %s\n", ourProject.id, ourProject.name, ourProject.description)
		}

	},
}

func init() {
	listCmd.AddCommand(listProjectsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listProjectsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listProjectsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func searchForProject(db *sql.DB) []project {
	rows, err := db.Query("SELECT id, name, description FROM projects")
	defer rows.Close()
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	projects := []project{}
	for rows.Next() {
		ourProject := project{}
		err = rows.Scan(&ourProject.id, &ourProject.name, &ourProject.description)
		if err != nil {
			log.Fatal(err)
		}

		projects = append(projects, ourProject)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return projects
}
