package main

import (
	"fmt"
	"time"
)

func main() {
	//CONCORRENCIA != PARALELISMO
	//PARALELISMO executa tarefas em paralelo se um processador tiver mais um nucleo
	//CONCORRENCIA não necessariamente executa as tarefas ao mesmo tempo, se tiver
	//somente um nucleo ele vai executar um pouco de cada tarefa(revesamento)

	//se executarmos dessa forma ele só vai executar a primeira função quando a segunda
	//acabar, porém elas são funções que nao param de funcionar, então a segunda nunca
	//vai ser executada dessa forma
	// escrever("Hellow World")
	// escrever("Golang")

	//utilizamos o comando go para executar a segunda função sem esperar a primeira
	//se utilizarmos go na segunda função ele nao vai printar nada, pois ele entendeu
	//que é pra executar as duas sem esperar nada, aí a aplicação termina antes de
	//poder executar as duas funções
	go escrever("Hello World")
	escrever("Golang")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
