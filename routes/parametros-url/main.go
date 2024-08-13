package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// O path com {} indica que eh um parametro, pode ser combinado
	// com metodos como "POST /alunos/{turma}/{id}" ou "DELETE /alunos/{turma}/{id}"
	mux.HandleFunc("/alunos/{turma}/{id}", Handlerzinho)

	http.ListenAndServe(":80", mux)
}

func Handlerzinho(w http.ResponseWriter, r *http.Request) {
	// Acessando os parametros da URL usando o metodo PathValue, que
	// busca no mapa de parametros o valor correspondente a chave passada   ↓↓.
	fmt.Fprintf(w, "Parametros: %v, %v", r.PathValue("turma"), r.PathValue("id"))
}
