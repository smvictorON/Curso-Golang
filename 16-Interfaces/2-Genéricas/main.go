package main

import "fmt"

//o tipo desse parametro é uma interface sem nenhuma assinatura
//ou seja generica, qualquer coisa atende essa interface, como
//ela atende todo tipo, o tipo realmente fica desnecessario, porém
//o golang é uma linguagem fortemente tipada o que vai contra isso
//tem que ser usado somente quando necessario para nao virar gambiarra
//um exemplo é o Println que recebe varias interfaces, qualquer coisa
//só pra ser printado
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("Victor")
	generica(28)
	generica(true)
	generica(123.45)

	//um map totalmente genérico e frouxo, ele aceita qualquer coisa
	//semelhante a um objeto no js
	mapa := map[interface{}]interface{}{
		1:            "string",
		float32(100): true,
		"string":     10,
	}

	fmt.Println(mapa)
}
