package main

import "fmt"

func main() {
	//map é uma estrutura que podemos guardar dados sendo
	//o primeiro parametro o tipo das chave e o segundo o tipo dos valores
	user := map[string]string{
		"nome":      "Victor",
		"sobrenome": "Mussio",
	}

	fmt.Println(user)
	fmt.Println(user["nome"])

	//aninhando maps
	user2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Victor",
			"ultimo":   "Mussio",
		},
		"curso": {
			"nome":      "ADS",
			"faculdade": "Fatec",
		},
	}

	fmt.Println("-------------")
	fmt.Println(user2)
	fmt.Println(user2["nome"])
	fmt.Println(user2["nome"]["primeiro"])

	//usamos o delete para remover campos do map
	delete(user2["curso"], "faculdade")

	fmt.Println("-------------")
	fmt.Println(user2)
	delete(user2, "curso")

	fmt.Println("-------------")
	fmt.Println(user2)

	//adicionar um campo no map após criado
	user2["conjuge"] = map[string]string{
		"nome": "Lamara",
	}

	fmt.Println("-------------")
	fmt.Println(user2)

}
