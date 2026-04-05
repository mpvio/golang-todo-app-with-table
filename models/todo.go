package models

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

var idInc AutoIncID

type Todo struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

// slice to hold multiple todos
type Todos []Todo

// functions for slice manipulation
func (todos *Todos) Add(title string) {
	newTodo := Todo{
		ID:        idInc.Next(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	*todos = append(*todos, newTodo)
}

func (todos *Todos) Get(id int) *Todo {
	t := *todos
	for i := range t {
		if t[i].ID == id {
			return &t[i]
		}
	}
	return nil
}

func (todos *Todos) Delete(id int) error {
	t := *todos
	for i := range t {
		if t[i].ID == id {
			*todos = append(t[:i], t[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}

func (todos *Todos) Edit(id int, title string) error {
	t := *todos
	for i := range t {
		if t[i].ID == id {
			t[i].Title = title
			*todos = t
			return nil
		}
	}
	return errors.New("todo not found")
}

func (todos *Todos) Complete(id int) error {
	t := *todos
	for i := range t {
		if t[i].ID == id {
			if t[i].Completed {
				return errors.New("todo already completed")
			}
			t[i].Completed = true
			now := time.Now()
			t[i].CompletedAt = &now
			*todos = t
			return nil
		}
	}
	return errors.New("todo not found")
}

func (todos *Todos) Print() {
	t := *todos
	table := table.New(os.Stdout)
	table.SetHeaders("Id", "Title", "Completed", "Created At", "Completed At")
	table.SetAutoMerge(true)

	var completed string
	var completedAt string

	for _, todo := range t {
		if todo.Completed {
			completed = "Yes"
			completedAt = todo.CompletedAt.Format(time.RFC1123)
		} else {
			completed = "No"
			completedAt = ""
		}

		table.AddRow(
			strconv.Itoa(todo.ID),
			todo.Title,
			completed,
			todo.CreatedAt.Format(time.RFC1123),
			completedAt,
		)
	}

	table.Render()
}

// set ID counter to max ID in todos to prevent ID collisions when adding new todos
func (todos *Todos) SetAutoIncID() {
	highest := 0
	for _, todo := range *todos {
		if todo.ID > highest {
			highest = todo.ID
		}
	}
	idInc.Set(highest + 1)
}
