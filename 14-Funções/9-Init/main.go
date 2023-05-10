package main

import "fmt"

var n int

//funções init rodam antes da função main, usada para iniciar variaveis
func init() {
	n = 10
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
	fmt.Println(n)
}
