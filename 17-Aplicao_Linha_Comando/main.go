package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	aplicacacao := app.Gerar()

	if erro := aplicacacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
