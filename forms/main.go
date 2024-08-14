package main

import (
	"log"
	"net/http"
)

type Formulario struct {
	Nome     string
	Pergunta string
}

func main() {
	// Criamos um mux para registrar as rotas.
	mux := http.NewServeMux()
	// GET para exibir o formulario.
	mux.HandleFunc("GET /pergunta", Pergunta)
	// POST para enviar o formulario.
	mux.HandleFunc("POST /pergunta", NovaPergunta)

	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatalf("Erro ao iniciar: %v", err)
	}
}

func NovaPergunta(w http.ResponseWriter, r *http.Request) {
	// r.FormValue("nome") pega o valor do campo nome, os valores
	// dos campos serao armazenados na struct Formulario.
	resposta := Formulario{
		Nome:     r.FormValue("nome"),
		Pergunta: r.FormValue("pergunta"),
	}

	log.Printf("Nova pergunta\n Nome: %s\n Pergunta: %s", resposta.Nome, resposta.Pergunta)
	// Redireciona para a pagina de pergunta, nao mostra nenhum status
	// ou confirmacao para o usuario saber se houve sucesso ou nao.
	http.Redirect(w, r, "/pergunta", http.StatusSeeOther)
}

func Pergunta(w http.ResponseWriter, r *http.Request) {
	// Apenas serve o arquivo como formulario.
	http.ServeFile(w, r, "pergunta.html")
}
