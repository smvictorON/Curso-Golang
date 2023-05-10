package main

import (
	"fmt"
)

func main() {
	// var/const funcionam da mesma forma, porém constantes nao podem ser modificadas

	//declaração explicita, dizendo o tipo
	var str1 string = "String 1"
	fmt.Println(str1)

	//declaracao implicita ele define pelo valor que recebe (inferencia de tipo)
	str2 := "String 2"
	fmt.Println(str2)

	//declaração explicita para mais de uma variável
	var (
		str3 string = "String 3"
		str4 string = "String 4"
	)
	fmt.Println(str3, str4)

	//declaração implicita para mais de uma variável
	str5, str6 := "String 5", "String 6"
	fmt.Println(str5, str6)

	//invertendo valores das variáveis
	str5, str6 = str6, str5
	fmt.Println(str5, str6)
}
