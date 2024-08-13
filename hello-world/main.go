package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a rota "/" e associa a função HandlerHello a ela
	http.HandleFunc("/", HandlerHello)

	// Comeca a servir na porta 80 e aguarda por requests
	// (Portas abaixo de 1024 requerem permissao de su)
	http.ListenAndServe(":80", nil)
}

func HandlerHello(w http.ResponseWriter, r *http.Request) {
	// Escreve no writer `w`, usa o request `r` para obter informacoes do request
	fmt.Fprintf(w, "Path do request: %s\nMétodo usado: %s\n", r.URL.Path, r.Method)
}
