package main

import "fmt"

//parametro por copia
func inverter(num int) int {
	return num * -1
}

//como estamos alterando direto no endereço, nao precisamos de retorno
//parametro por referencia
func inverterPont(num *int) {
	*num = *num * -1
}

func main() {
	numero := 20
	numeroInvertido := inverter(numero)
	fmt.Println(numeroInvertido)
	//numero continua inauterado
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	//dessa forma ele altera o valor original
	inverterPont(&novoNumero)
	fmt.Println(novoNumero)
}
