package funcoes

import "fmt"


func FuncaoVariaticaComMaisDeUmParametro(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	fmt.Println("FuncaoVariaticaComMaisDeUmParametro: ", total)
	return total
}

func FuncaoVariaticaComMaisDeUmParametroComRetorno(texto string, numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	fmt.Println("FuncaoVariaticaComMaisDeUmParametroComRetorno: ", texto, total)
	return total
}