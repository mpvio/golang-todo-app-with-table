package main

import (
	"fmt"
	"todo-app-table/models"
)

func main() {
	// create todos model
	var todos models.Todos
	// create storage for todos
	storage := models.NewStorage[models.Todos]("todos.json")
	// load todos & set auto increment ID
	storage.Load(&todos)
	todos.SetAutoIncID()
	// print todos to console
	todos.Print()
	fmt.Println("Use -h flag for command options.")
	// create command flags and execute command
	cmdFlags := models.NewCmdFlags()
	cmdFlags.Execute(&todos)
	// save todos to file
	storage.Save(todos)
}
