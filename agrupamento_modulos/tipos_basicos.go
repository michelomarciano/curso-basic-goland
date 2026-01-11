package agrupamento_modulos

import (
	"fmt"
	"modulo/tiposdedados"
	"modulo/variaveis"
)

func ExecutarTiposBasicos() {
	fmt.Println("--------------------------------")
	fmt.Println("VARIAVEIS")
	fmt.Println("--------------------------------")
	fmt.Println("VARIAVEL IMPLICITA:")
	variaveis.VariavelImplicita()
	fmt.Print("\n")
	fmt.Println("VARIAVEL EXPLICITA:")
	variaveis.VariavelExplicita()
	fmt.Print("\n")

	fmt.Println("--------------------------------")
	fmt.Println("TIPOS DE DADOS")
	fmt.Println("--------------------------------")
	fmt.Println("INT:", tiposdedados.Int())
	fmt.Println("UINT:", tiposdedados.Uint())
	fmt.Println("FLOAT:", tiposdedados.Float())
	fmt.Println("CHAR:", tiposdedados.Char())
	fmt.Println("BOOL:", tiposdedados.Bool())
	fmt.Println("ERRO:", tiposdedados.Erro())
	fmt.Print("\n")
}

