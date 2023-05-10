package main

import "fmt"

//toda função recursiva deve ter um ponto de parada senão
//pode causar o que é chamado de estouro de pilha(stack overflow)
func fibonacci(pos uint) uint {
	if pos <= 1 {
		return pos
	}
	return fibonacci(pos-2) + fibonacci(pos-1)
}

func reduzir(num uint) uint {
	if num <= 1 {
		return num
	}
	return reduzir(num - 1)
}

func main() {
	pos := uint(10)

	for i := uint(1); i < pos; i++ {
		fmt.Println(fibonacci(i))
	}
	fmt.Println(fibonacci(pos))
	fmt.Println(reduzir(pos))
}
