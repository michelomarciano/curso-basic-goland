package funcoes

import "fmt"

func FuncaoClosure() func() {
	texto := "Dentro da funcao closure"
	var funcao = func() {
		fmt.Println(texto)
	}
	return funcao
}
