package main

import (
	"fmt"
	"time"
)

//a vantagem desse padrão é que nao função main nao precisamos no preocupar em criar
//canais e goroutines, isso tudo é a função escrever quem cuida
func main() {
	canal := escrever("Hello World")

	//conforme o canal vai sendo alimentado na função escrever, aqui vai imprimindo
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
