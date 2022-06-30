/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/in2tivetech/ozone-ce/controllers"
	"github.com/spf13/cobra"
)

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
		fmt.Println("createProject called with arguments ", args)
		fmt.Println("the project name is : ", args[0])
		controllers.CreateProject(args[0], args[1])
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

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
