package funcoes

import "fmt"

func Defer() {
	fmt.Println("Funcao com Defer")
}

func SemDefer() {
	fmt.Println("Funcao Sem Defer")
}

func AlunoAprovado(n1, n2 float64) bool {
	defer fmt.Println("Media calculada. Resultado será retornado")
	fmt.Println("Calculando media...")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}