package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	//quantidade de funções que vão ser executadas
	waitGroup.Add(2)

	go func() {
		escrever("Hello World")
		//sinaliza que terminou
		waitGroup.Done() // -1
	}()
	go func() {
		escrever("Golang")
		waitGroup.Done() // -1
	}()

	//o wait espera as duas goroutines acabarem para terminar, se remover o wait
	//acontece igual o da primeira aula e a aplicação termina antes de executar as
	//goroutines
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
