package ponteiro

import (
	"fmt"
)
// Nesse exemplo a variavel1 é uma cópia da variavel2, ou seja, se alterarmos o valor da variavel1, o valor da variavel2 não será alterado.
func AtribuiValorParaVariavel() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println("Valor da variavel1:", variavel1, "Valor da variavel2:", variavel2)

    variavel1++
    fmt.Println("Valor da variavel1:", variavel1, "Valor da variavel2:", variavel2)
}

// Nesse exemplo o ponteiro é uma referência para a variavel1, ou seja, se alterarmos o valor da variavel1, o valor do ponteiro também será alterado.
func DiferencaEntrePonteiroEValor() {
    var variavel1 int
    var ponteiro *int

    variavel1 = 10
    ponteiro = &variavel1

    fmt.Println("DiferencaEntrePonteiroEValor - Valor da variavel1:", variavel1, "Valor do ponteiro:", *ponteiro)

    variavel1++
    fmt.Println("DiferencaEntrePonteiroEValor - Valor da variavel1:", variavel1, "Valor do ponteiro:", *ponteiro)
}


// Nesse exemplo o ponteiro é uma referência para a variavel1, ou seja, se alterarmos o valor da variavel1, o valor do ponteiro também será alterado.
func ModificarValorDaVariavelApontadaPorPonteiro() {
    var variavel1 int
    var ponteiro *int

    variavel1 = 11
    ponteiro = &variavel1

    fmt.Println("ModificarValorApontadoPorPonteiro - Valor da variavel1:", variavel1, "Valor do ponteiro:", *ponteiro, "Endereço do ponteiro:", ponteiro)
}

