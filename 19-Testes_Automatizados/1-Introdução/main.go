package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereço("Rod Paulista")
	fmt.Println(tipoEndereco)
}
