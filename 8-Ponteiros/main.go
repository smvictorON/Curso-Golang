package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println("-------")

	var1++
	var2--

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println("-------")

	var var3 int
	var pont3 *int

	fmt.Println(var3)
	fmt.Println(pont3)
	fmt.Println("-------")

	var3 = 100
	pont3 = &var3

	fmt.Println(var3)
	fmt.Println(pont3)
	fmt.Println(*pont3)
	fmt.Println("-------")

	var3 = 150

	fmt.Println(var3)
	fmt.Println(*pont3)
	fmt.Println("-------")
}
