package main

import "fmt"

//dessa forma podemos passar quantos parametros quiser
//só podemos ter um parametro nesse padrao e deve ser o ultimo
func soma(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

//usando os 2 tipos de parametros
func mixado(label string, numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	fmt.Println(label)
	return total
}

func main() {
	res := soma(1, 2, 3, 4)
	fmt.Println(res)
	res = soma(1, 2, 3, 4, 5, 6)
	fmt.Println(res)
	fmt.Println(res)
	res = mixado("Aqui", 1, 2)
	fmt.Println(res)
}
