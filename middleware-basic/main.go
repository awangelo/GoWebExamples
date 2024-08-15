package main

import (
	"fmt"
	"log"
	"net/http"
)

// logging sera o middleware simples, ira fazer alguma
// tarefa e repessar o request para o seu handler.
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Novo request para:", r.URL.Path)
		// Repassa o request para o handler
		f(w, r)
	}
}

func main() {
	http.HandleFunc("GET /login", logging(login))

	http.ListenAndServe(":80", nil)
}

// Essa funcao ainda sera um handler comum.
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Pagina de login</h1>")
}
