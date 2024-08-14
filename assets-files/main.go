package main

import "net/http"

func main() {
	// FileServer retorna um handler que serve arquivos de um sistema de arquivos.
	fs := http.FileServer(http.Dir("static/"))
	// StripPrefix retorna um handler que serve o conteudo de um handler sem o prefixo especificado.
	// ou seja, o prefixo "/static/" é removido da URL para que o arquivo seja encontrado.
	// E entao, Handle registra o handler para o padrão de URL especificado.
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", Htmlzinho)

	http.ListenAndServe(":80", nil)
}

func Htmlzinho(w http.ResponseWriter, r *http.Request) {
	// ServeFile responde ao request HTTP enviando o conteudo do arquivo ou diretorio.
	http.ServeFile(w, r, "index.html")
}
