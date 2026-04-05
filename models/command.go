package models

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type CmdFlags struct {
	Add    string
	Delete int
	Edit   int
	// EditTitle string
	Complete int
	Print    bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo with the given title")
	flag.IntVar(&cf.Delete, "delete", -1, "Delete the todo with the given ID")
	flag.IntVar(&cf.Edit, "edit", -1, "Edit the todo with the given ID")
	// flag.StringVar(&cf.EditTitle, "edit-title", "", "Edit the title of the todo with the given ID")
	flag.IntVar(&cf.Complete, "complete", -1, "Complete the todo with the given ID")
	flag.BoolVar(&cf.Print, "print", false, "Print all todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Delete != -1:
		todos.Delete(cf.Delete)
	case cf.Complete != -1:
		todos.Complete(cf.Complete)
	case cf.Print:
		todos.Print()
	case cf.Edit != -1:
		var newTitle string
		fmt.Printf("Enter new title for todo with ID %d: ", cf.Edit)

		reader := bufio.NewReader(os.Stdin)
		newTitle, _ = reader.ReadString('\n')  // read input until (and including) newline
		newTitle = strings.TrimSpace(newTitle) // also trims newline

		todos.Edit(cf.Edit, newTitle)
		// default:
		// 	fmt.Println("No valid command provided.")
	}
}
