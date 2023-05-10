package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculos(n1, n2 int) (int, int, int) {
	soma := n1 + n2
	subt := n1 - n2
	mult := n1 * n2

	return soma, subt, mult
}

func main() {
	res := somar(1, 2)
	fmt.Println(res)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Texto")

	res1, res2, _ := calculos(5, 10)
	fmt.Println(res1, res2)
}
