package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidores"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "uol.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ips",
			Usage:  "Busca ips",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca servidores",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
