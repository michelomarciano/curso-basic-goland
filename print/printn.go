package print

import "fmt"

func Println() {
	var(
		nome string = "Mike"
        sobrenome string = "Marciano"
		salario float32 = 100.20
	)
	fmt.Println("Println: O nome é: " + nome + " sobrenome: " + sobrenome + " salario: " + fmt.Sprint(salario))
}