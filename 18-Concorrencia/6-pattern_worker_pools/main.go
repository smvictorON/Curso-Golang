package main

import "fmt"

func fibonacci(pos int) int {
	if pos <= 1 {
		return pos
	}
	return fibonacci(pos-2) + fibonacci(pos-1)
}

//as setas nos canais identifica qual tipo de canal ele é, no caso do primeiro
//ele só recebe dados e no segundo ele só envia dados
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	//aqui estamos chamando paralelamente o worker sem popular o canal de tarefas
	//dessa forma podemos chamar varias vezes o worker para ir pegando os valores
	//do canal da tarefa deixando mais rapido a execução, porém é limitado pelo
	//processador então nao adianta chamar mil vezes o worker
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		//conforme for populando as tarefas aqui ele vai recebendo no worker
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}
