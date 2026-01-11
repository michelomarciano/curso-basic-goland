package funcoes

import "fmt"

func FuncaoPonteiro(numero *int) {
	fmt.Println("FuncaoPonteiro", *numero)
	*numero = *numero * -1
}
