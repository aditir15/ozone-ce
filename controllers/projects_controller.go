package controllers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/in2tivetech/ozone-ce/models"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	// Connect to database
	db, err = sql.Open("sqlite3", "./names.db")
	checkErr(err)
}

func CreateProject(name, desc string) {

	// defer close
	// defer db.Close()

	newProject := models.Project{
		Name:        name,
		Description: desc,
	}
	//add a record
	addProject(db, newProject)
}

func ListProjects() {
	//read a record
	fmt.Println("Reading Records.......... ")
	project := searchForProject(db)
	for _, ourProject := range project {
		fmt.Printf("\n----\nID: %d\nName: %s\nDescription: %s\n", ourProject.Id, ourProject.Name, ourProject.Description)
	}
}

func addProject(db *sql.DB, newProject models.Project) {

	stmt, err := db.Prepare("INSERT INTO projects (id, name, description) VALUES (?, ?, ?)")
	checkErr(err)
	result, err := stmt.Exec(nil, newProject.Name, newProject.Description)
	checkErr(err)
	fmt.Printf("result %v", result)
	defer stmt.Close()

	fmt.Printf("Added %v %v \n", newProject.Name, newProject.Description)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func searchForProject(db *sql.DB) []models.Project {
	rows, err := db.Query("SELECT id, name, description FROM projects")
	defer rows.Close()
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	projects := []models.Project{}
	for rows.Next() {
		ourProject := models.Project{}
		err = rows.Scan(&ourProject.Id, &ourProject.Name, &ourProject.Description)
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
