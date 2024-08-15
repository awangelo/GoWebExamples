package main

import (
	"fmt"
	"log"
	"net/http"
)

// Middleware sera um tipo de funcao que recebe um handler e retorna outro handler.
type Middleware func(f http.HandlerFunc) http.HandlerFunc

// Novo middleware que ira logar o request e passar para o proximo middleware ou handler.
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Fazendo coisas de middleware...
			log.Println("Novo request para:", r.URL.Path)
			// Repassa o request para o handler
			f(w, r)
		}

	}
}

// Novo middleware que ira fazer algo dependendo do status code da resposta e passar para o proximo.
func StatusCode() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Fazendo coisas de middleware...
			log.Println("Fazendo algo dependendo do status code...")
			// Passa a funcao para o proximo middleware ou handler.
			f(w, r)
		}
	}
}

func main() {
	// Registra o handler com a funcao Chain.
	http.HandleFunc("/auth", Chain(auth, Logging(), StatusCode()))

	http.ListenAndServe(":80", nil)
}

// Sera um handler comum.
func auth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Auth</h1>")
}

// Recebe um handler e uma lista de middlewares e retorna um handler.
// Chain ira chamar cada middleware passando o handler como argumento,
// ou seja, cada funcao middleware que for passada sera chamada e passara
// o handler para a proxima funcao.
func Chain(f http.HandlerFunc, mws ...Middleware) http.HandlerFunc {
	for _, m := range mws {
		f = m(f)
	}
	return f
}
