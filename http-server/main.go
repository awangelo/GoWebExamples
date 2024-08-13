package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a rota "/" e associa a funcao HandlerHello a ela
	http.HandleFunc("/", HttpServer)

	// fs serve arquivos estaticos do diretorio "static/"
	fs := http.FileServer(http.Dir("static/"))
	// Metodos `Handle` e `HandleFunc` sao usados para associar um URL path a um handler
	// HandleFunc associa um URL path a uma funcao que implementa a interface `http.Handler`
	//    que tem assinatura: (http.ResponseWriter, *http.Request). Igual a linha 27.
	// Handle nesse caso associa o path "/static/" ao FileServer `fs` que serve arquivos estaticos
	//    e utiliza o metodo `StripPrefix` para remover o prefixo "/static/" do path do arquivo
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Inicia o servidor HTTP na porta 80
	http.ListenAndServe(":80", nil)
}

// Implementa a interface `http.Handler` que tem assinatura: (http.ResponseWriter, *http.Request)
func HttpServer(w http.ResponseWriter, r *http.Request) {
	gifPath := "/static/spin.gif"

	fmt.Fprintf(w, `
		<!DOCTYPE html>
			<img src="%s">
		</html>
	`, gifPath)
}
