package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

//se nao quiser que algum dos campos apareça no json bastas substituir
//o nome para - Ex: `json:"-"`
type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	//Marshal = map -> json ou struct -> json
	//Unmarshal = json -> map ou json -> struct
	//bytes.NewBuffer transforma um slice de bytes em um json para leitura

	c1 := cachorro{"Rex", "Dalmata", 10}
	fmt.Println(c1)

	cJson1, err := json.Marshal(c1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cJson1)
	fmt.Println(bytes.NewBuffer(cJson1))

	c2 := map[string]string{
		"nome": "Totó",
		"raça": "Poodle",
	}

	cJson2, err := json.Marshal(c2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cJson2)
	fmt.Println(bytes.NewBuffer(cJson2))

	cJson3 := `{ "nome": "Leo", "raca": "Pintcher", "idade": 8 }`

	var c3 cachorro

	erro := json.Unmarshal([]byte(cJson3), &c3)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c3)

	cJson4 := `{ "nome": "Pingo", "raca": "Shitzu"}`

	c4 := make(map[string]string)

	error := json.Unmarshal([]byte(cJson4), &c4)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(c4)
}
