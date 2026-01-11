package app

import (
	"log"
	"net"
	"fmt"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicacao de linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicacao de Linha de Comando"
	app.Usage = "Busca Ips e Nomes de Servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "mikemarciano.dev.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca Ips de endereco na internet",
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name: "servidores",
			Usage: "Busca o nome do servidor na internet",
			Flags: flags,
			Action: buscarServidor,
		},
	}

	return app
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

func buscarServidor(c *cli.Context){
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}