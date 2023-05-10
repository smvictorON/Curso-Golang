package main

import "fmt"

type user struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint16
}

func main() {
	var u user
	fmt.Println(u)

	u.nome = "Victor"
	u.idade = 28
	fmt.Println(u)

	end := endereco{"Antonio Castilho", 470}
	u2 := user{
		"Lamara",
		34,
		end,
	}
	fmt.Println(u2)

	//quando queremos passar só um campo da struct, temos que declarar qual é
	u3 := user{nome: "Gamba"}
	u4 := user{idade: 10}
	fmt.Println(u3)
	fmt.Println(u4)
}
