package ifelse

import "fmt"

func IfElse() {
	var idade int = 18

	if idade >= 18 {
		fmt.Println("Você é maior de idade")
	} else {
		fmt.Println("Você é menor de idade")
	}
}

func IfElseInicializandoVariavel() {
	var numero int = 12

	if idade := numero; idade >= 18 {
		fmt.Println("Você é maior de idade")
	} else {
		fmt.Println("Você é menor de idade")
	}
}
