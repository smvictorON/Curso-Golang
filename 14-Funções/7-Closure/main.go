package main

import "fmt"

//funçao closure tem como se fosse uma memória e pode printar
//dados do escopo inicial em outras funções
func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcNova := closure()

	funcNova()
}
