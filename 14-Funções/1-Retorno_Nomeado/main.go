package main

import "fmt"

func calculos(n1, n2 int) (soma int, subtracao int) {
	//como o retorno é nomeado nao precisa usar o :=
	//e tambem não precisa retorna diretamente elas
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	res1, res2 := calculos(1, 2)

	fmt.Println(res1, res2)
}
