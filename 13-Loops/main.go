package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// 	//função semelhante ao settimeout
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)

	// //j só pode ser acessada dentro do for
	// for j := 0; j < 10; j += 2 {
	// 	fmt.Println(j)
	// 	time.Sleep(time.Second)
	// }

	// nomes := [3]string{"Victor", "Lamara", "Laura"}

	// for indice, nome := range nomes {
	// 	fmt.Println(indice)
	// 	fmt.Println(nome)
	// 	time.Sleep(time.Second)
	// }

	// //iterando por letra, retorna o valor asc da letra se usar o range
	// for indice, letra := range "PALAVRA" {
	// 	fmt.Println(indice)
	// 	fmt.Println(letra)
	// 	fmt.Println(string(letra))
	// 	time.Sleep(time.Second)
	// }

	// users := map[string]string{
	// 	"nome":      "Victor",
	// 	"sobrenome": "Mussio",
	// }

	// for chave, valor := range users {
	// 	fmt.Println("Chave:", chave)
	// 	fmt.Println("Valor:", valor)
	// 	time.Sleep(time.Second)
	// }

	for {
		fmt.Println("Loop Infinito")
		time.Sleep(time.Second)
	}
}
