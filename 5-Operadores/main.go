package main

import "fmt"

func main() {
	//aritmeticos
	soma := 10 + 20
	subtracao := 10 - 20
	multiplicacao := 10 * 20
	divisao := 20 / 10
	restoDaDivisao := 20 % 6

	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)

	var n1 int16 = 10
	var n2 int32 = 20
	// nao tem como fazer essa conta dessa forma
	// resSoma := n1 + n2
	// forma correta de se fazer essa conta
	resSoma := n1 + int16(n2)
	fmt.Println(resSoma)

	//atribuição, pode ser usado = ou :=
	var ex1 string = "text"
	ex2 := "text"
	fmt.Println(ex1, ex2)

	//relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//logicos && || !
	fmt.Println("-----------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//unários
	num := 10
	num++
	fmt.Println(num)
	num += 15
	fmt.Println(num)
	num--
	fmt.Println(num)
	num -= 5
	fmt.Println(num)
	num *= 3
	fmt.Println(num)
	num /= 2
	fmt.Println(num)
	num %= 4
	fmt.Println(num)

}
