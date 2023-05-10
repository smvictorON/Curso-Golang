package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr1 [5]int
	arr1[0] = 1
	fmt.Println(arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	// dessa forma conseguimos colocar a quantidade de itens sem declarar
	// mas após isso ele vai conter o tamanho do total dos itens
	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr3)

	slice := []int{1, 2, 3}
	fmt.Println(slice)

	slice = append(slice, 1)
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(arr3))
	fmt.Println(reflect.TypeOf(slice))

	// nesse caso ele pega o 2 e o 3, o primeiro parametro conta e o segundo não
	slice2 := arr2[2:4]
	fmt.Println(slice2)

	arr2[2] = 30
	fmt.Println(arr2)
	fmt.Println(slice2)

	//array interno
	fmt.Println("-------")
	//make(tipo, tamanho, capacidade)
	//quando omitido a capacidade, ela vai ser igual ao tamanho
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	//tamanho
	fmt.Println(len(slice3))
	//capacidade
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 1)
	fmt.Println(slice3)
	slice3 = append(slice3, 2)
	fmt.Println(slice3)

	//quando estoura um slice ele aumenta a capacidade do slice
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	fmt.Println("-------")
	slice4 := make([]float32, 5)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
