package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoList struct {
	Name  string
	Todos []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoList{
			Name: "My TODO list",
			Todos: []Todo{
				{Title: "Install Go", Done: true},
				{Title: "Learn Go", Done: false},
				{Title: "Build Website", Done: false},
			},
		}
		tmpl.Execute(w, data)
	})

	log.Panic(http.ListenAndServe(":8080", nil))
}
