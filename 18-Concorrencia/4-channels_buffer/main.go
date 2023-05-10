package main

import (
	"fmt"
)

func main() {
	//o segundo parametro aqui é o buffer, que é a capacidade, ele só vai bloquear
	//quando atingir a capacidade, e se nao tiver buffer ele bloqueia sempre
	canal := make(chan string, 2)

	//executando aqui ele vai ficar bloqueado e nunca vai poder receber esse valor
	canal <- "Hello World"
	canal <- "Golang"

	msg := <-canal
	msg2 := <-canal

	fmt.Println(msg)
	fmt.Println(msg2)

}
