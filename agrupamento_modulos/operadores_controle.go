package agrupamento_modulos

import (
	"fmt"
	"modulo/ifelse"
	"modulo/loops"
	"modulo/operadores"
	"modulo/switchs"
)

func ExecutarOperadores() {
	fmt.Println("--------------------------------")
	fmt.Println("OPERADORES")
	fmt.Println("--------------------------------")
	fmt.Print("OPERADORES ARITMETICOS: ")
	operadores.OperadoresAritmeticos()
	fmt.Print("OPERADORES RELACIONAIS: ")
	operadores.OperadoresRelacionais()
	fmt.Print("OPERADORES LOGICOS: ")
	operadores.OperadoresLogicos()
	fmt.Print("OPERADORES LOGICOS COM 3 COMBINACOES: ")
	operadores.OperadoresLogicosTresCombinacoes()
	fmt.Print("\n")
}

func ExecutarControleFluxo() {
	fmt.Println("--------------------------------")
	fmt.Println("IF ELSE")
	fmt.Println("--------------------------------")
	ifelse.IfElse()
	ifelse.IfElseInicializandoVariavel()
	fmt.Print("\n")

	fmt.Println("--------------------------------")
	fmt.Println("SWITCH")
	fmt.Println("--------------------------------")
	switchs.Switch()
	fmt.Println(switchs.SwitchComRetorno())
	fmt.Println(switchs.SwitcComAtribuicaoDeVariavel(1))
	fmt.Print("\n")

	fmt.Println("--------------------------------")
	fmt.Println("LOOPS")
	fmt.Println("--------------------------------")
	loops.LoopFor()
	loops.LoopWhile()
	loops.LoopDoWhile()
	loops.LoopForRange()
	loops.LoopForRangeString()
	loops.LoopForRangeMap()
	fmt.Print("\n")
}

