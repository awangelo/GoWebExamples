package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Mux eh um "roteador", ou seja, ele recebe uma requisicao e
	// direciona para a funcao correta (um handler).
	mux := http.NewServeMux()
	// O metodo HandleFunc recebe um padrao de URL e direciona para um handler.
	mux.HandleFunc("GET /teste", Get)
	mux.HandleFunc("POST /teste", Post)

	// O ListenAndServe recebe um endereco e um handler.
	// O endereco eh o endereco que o servidor vai escutar.
	// O handler eh o "roteador" que vai direcionar a requisicao para a funcao correta.
	http.ListenAndServe(":80", mux)
}

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Voce usou get: ", r.Method)
}

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.Method)
}
