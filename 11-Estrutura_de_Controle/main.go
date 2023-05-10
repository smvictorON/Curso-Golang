package main

import "fmt"

func main() {
	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	//if init declara a var em tempo de execucao para a checagem
	//essa var nao pode ser acessada fora do escopo
	if numero2 := numero; numero2 > 0 {
		fmt.Println(numero2)
		fmt.Println("Maior que 0")
	} else if numero2 < -10 {
		fmt.Println("Menor que 0")
	} else {
		fmt.Println("Entro 0 e -10")
	}
}
