package main

import "fmt"

func main() {

	res := func(texto string) string {
		return fmt.Sprintf("Recebido: %s %d", texto, 10)
	}("Função Anônima com Parâmetro e Retorno")

	fmt.Println(res)
}
