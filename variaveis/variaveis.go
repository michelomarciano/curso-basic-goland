package variaveis

import "fmt"

func VariavelExplicita() {

	// declarando uma variavel explicita
	var nome string = "Mike"
	fmt.Println(nome)

	// declarando mais de uma variavel ao mesmo tempo
	var (
		nome1      string = "Mike"
		sobrenome1 string = "marciano"
	)
	fmt.Printf("Nome: %s, Sobrenome: %s\n", nome1, sobrenome1)
}

func VariavelImplicita() {

	// declarando mais de uma variavel ao mesmo tempo
	nome2, sobrenome2 := "Mike", "Marciano"
	fmt.Println(nome2, sobrenome2)

	// declarando uma variavel
	sobrenome := "marciano"
	fmt.Printf("O sobrenome é %s", sobrenome)

}
