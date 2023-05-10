package main

import (
	"log"
	"net/http"
)

func users(w http.ResponseWriter, r *http.Request) {
	//escreve uma mensagem com o http
	w.Write([]byte("Usuários"))
}

func main() {
	//http é um protocolo de comunicação
	//cliente(faz a req) - servidor(processa a req e devolve a res)
	//rotas = URI(identificador do recurso) + métodos(GET, POST, PUT, DELETE)

	//pode ser feito diretamente dessa maneira
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		//escreve uma mensagem com o http
		w.Write([]byte("Hello World"))
	})

	//ou com uma função
	http.HandleFunc("/users", users)

	//fica escutando a porta 5000 para receber as requisições
	log.Fatal(http.ListenAndServe(":5000", nil))

}
