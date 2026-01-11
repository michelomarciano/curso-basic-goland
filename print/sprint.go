package print

import "fmt"

func Sprint() {
	var(
		salario float32 = 100.20
	)

	valorSalario := fmt.Sprint(salario)
	fmt.Println("O salario é " + valorSalario)
	fmt.Println("O salario é ", salario)

}