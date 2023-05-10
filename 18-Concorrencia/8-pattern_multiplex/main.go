package main

import (
	"fmt"
	"math/rand"
	"time"
)

//a ideia desse padrão é unir mais de um canal em um só
func main() {
	canal := multiplexar(escrever("Golang1"), escrever("Golang2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

//recebe dois canais e conforme vai chegando mensagem nos canais ele jogo em um só
func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)
	go func() {
		for {
			select {
			case msg := <-canalEntrada1:
				canalSaida <- msg
			case msg := <-canalEntrada2:
				canalSaida <- msg
			}
		}
	}()

	return canalSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			//utilizando rand e time aqui podemos simular um tempo de espera
			//'aleatorio', porém segue um padrão após executar algumas vezes
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
