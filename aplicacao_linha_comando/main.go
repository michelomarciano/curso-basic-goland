package main

import (
	"log"
	"os"

	"modulo/aplicacao_linha_comando/app"
)

func main() {
	aplicacao := app.Gerar()
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

