package funcoes

import "fmt"

func FuncaoComRetorno(n1, n2 int8) string {
	return fmt.Sprintf("Funcao com retorno: %d + %d = %d", n1, n2, n1 + n2)
}

func RecuperandoValorDaFuncaoComVariavel(n1, n2 int8) int8 {
	soma := n1 + n2
	return soma
}

func FuncaoSemRetorno() {
	fmt.Println("Funcao sem retorno")
}

var PassandoFuncaoParaVariavel = func(n1, n2 int8) int8 {
	fmt.Println("Passando funcao para variavel")
	return n1 + n2
}





