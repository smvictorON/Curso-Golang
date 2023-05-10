package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{
		"Victor",
		"Mussio",
		28,
	}

	e1 := estudante{
		p1,
		"ADS",
		"Fatec",
	}

	fmt.Println(p1)
	fmt.Println(e1)
	fmt.Println(e1.nome)
}
