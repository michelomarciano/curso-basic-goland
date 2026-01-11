package print

import "fmt"

func Printf() {
	var(
		nome string = "Mike"
        sobrenome string = "Marciano"
		salario float32 = 100.20
	)
	fmt.Printf("\nNome %s e sobrenome %s", nome, sobrenome)
	fmt.Printf("\nSalario %f", salario)
}