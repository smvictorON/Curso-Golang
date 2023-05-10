package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	//dessa forma ele só vai printar após receber a mensagem no segundo canal que é
	//mais demorado do que o primeiro
	// for {
	// 	msg1 := <-canal1
	// 	fmt.Println(msg1)

	// 	msg2 := <-canal2
	// 	fmt.Println(msg2)
	// }

	//no go temos uma estrutura chamada select que é como se fosse um switch para canais
	//que isso melhora a performance porque ele nao espera mais um canal ser respondido
	//conforme cada um é respondido ele ja printa
	for {
		select {
		case msg1 := <-canal1:
			fmt.Println(msg1)
		case msg2 := <-canal2:
			fmt.Println(msg2)
		}
	}
}
