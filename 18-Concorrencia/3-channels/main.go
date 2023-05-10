package main

import (
	"fmt"
	"time"
)

func main() {
	//usamos channels para comunicar entre as goroutines e conseguir sincroniza-las
	//nesse caso um canal de strings
	canal := make(chan string)
	go escrever("Hello World", canal)

	//aqui estamos esperando recebendo o valor, ele para a execução esperando receber
	//todo envio e recebimento entre os canais são operações bloqueantes
	//nesse caso após ele receber a primeira vez ele vai seguir e finalizar a aplicacao
	// mensagem := <-canal

	//para melhorar isso colocamos dentro de um for, se não tiver um ponto de saida
	//após ele ser executando as 5x do escrever ele vai continuar esperando para receber
	//dando o famoso deadlock (all goroutines are asleep), tomar cuidado com isso porque
	//o go nao consegue identificar deadlocks no momento da compilação
	//o aberto le se o canal está fechado para poder sair evitando o deadlock
	// for {
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	//isso faz o mesmo que o for de cima sem precisar de uma var para verificar se o canal
	//está aberto, com o range ele ve se o canal está aberto, se nao tiver ele para
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		//aqui estamos mandando o valor texto para dentro do canal
		canal <- texto
		time.Sleep(time.Second)
	}

	//fecha o canal
	close(canal)
}
