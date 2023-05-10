package main

import "fmt"

func f1() {
	fmt.Println("Função 1")
}
func f2() {
	fmt.Println("Função 2")
}

//aqui temos um exemplo de uso, mesmo antes de calcular a média eu coloco
//o print com defer para ser impresso somente no final, ao invés de duplicar
//o print e colocalos antes dos 2 retornos, é bastante usado para executar algo
//independente do ponto de saida, Ex: fechar conexao com o banco
func alunoAprovado(n1, n2 int) bool {
	defer fmt.Println("Média foi calculada!")
	fmt.Println("Entrou")

	media := n1 + n2/2

	if media >= 6 {
		return true
	}

	return false
}

//defer adia a função para executar todas funções antes dela
//ela é executada antes da função terminar
func main() {
	// defer f1()
	// f2()

	res := alunoAprovado(6, 8)
	fmt.Println(res)
}
