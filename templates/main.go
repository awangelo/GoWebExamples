package main

import (
	"net/http"
	"text/template"
)

type Tarefa struct {
	Title string
	Done  bool
}

type TarefaPageData struct {
	PageTitle string
	Todos     []Tarefa
}

func main() {
	http.HandleFunc("/todo", Todo)
	http.ListenAndServe(":80", nil)
}

func Todo(w http.ResponseWriter, r *http.Request) {
	// data sera o objeto que sera passado para o template.
	data := TarefaPageData{
		PageTitle: "My TODO list",
		Todos: []Tarefa{
			{Title: "Tarefa 1", Done: false},
			{Title: "Tarefa 2", Done: true},
			{Title: "Tarefa 3", Done: true},
		},
	}

	// ParseFiles carrega o arquivo html e o transforma em um template.
	tmpl, err := template.ParseFiles("todo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute junta o template com o objeto data e escreve o resultado no writer.
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
